package parse

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	TsFileExtension        = ".ts"
	Mp4FileExtension       = ".mp4"
	M4FileExtensionPrefix  = ".m4"
	Mp4FileExtensionPrefix = ".mp4"
	CmfFileExtensionPrefix = ".cmf"
)

type ReadFile func(path string) ([]byte, error)

type Parser struct {
	ReadFile ReadFile
}

func NewParser() Parser {
	return Parser{
		ReadFile: func(path string) ([]byte, error) {
			return ioutil.ReadFile(path)
		},
	}
}

func (p *Parser) Parse(path string) (MasterPlaylist, error) {
	masterPlaylist, err := p.ParseMasterPlaylist(path)
	if err != nil {
		return MasterPlaylist{}, err
	}
	return masterPlaylist, nil
}

func (p *Parser) ParseMasterPlaylist(path string) (MasterPlaylist, error) {
	data, err := p.ReadFile(path)
	if err != nil {
		return MasterPlaylist{}, err
	}
	reader := bufio.NewReader(bytes.NewReader(data))
	var tags []string
	var mediaPlaylists []MediaPlaylist
	for {
		readBytes, _, err := reader.ReadLine()
		line := string(readBytes)
		if err == io.EOF {
			break
		}
		if err != nil {
			return MasterPlaylist{}, err
		}
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "# ") {
			continue
		}
		if strings.HasPrefix(line, "#EXT") {
			tags = append(tags, line)
		}
		if !strings.HasPrefix(line, "#") {
			mediaPlaylist, err := p.ParseMediaPlaylist(filepath.Join(filepath.Dir(path), line))
			if err != nil {
				return MasterPlaylist{}, err
			}
			mediaPlaylists = append(mediaPlaylists, mediaPlaylist)
		}
	}

	masterPlaylist := MasterPlaylist{
		Path:           path,
		Tags:           tags,
		MediaPlaylists: mediaPlaylists,
	}
	return masterPlaylist, nil
}

func (p *Parser) ParseMediaPlaylist(path string) (MediaPlaylist, error) {
	data, err := p.ReadFile(path)
	if err != nil {
		return MediaPlaylist{}, err
	}
	reader := bufio.NewReader(bytes.NewReader(data))
	var tags []string
	var segments []Segment
	var totalDurationMs float64
	var discontinuitySequence int
	for {
		readBytes, _, err := reader.ReadLine()
		line := string(readBytes)
		fmt.Println(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			return MediaPlaylist{}, err
		}
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "# ") {
			continue
		}
		// append EXT-X-MEDIA-SEQUENCE, EXT-X-DISCONTINUITY-SEQUENCE while ignoring pre-existed those tags.
		if strings.HasPrefix(line, "#EXT-X-MEDIA-SEQUENCE") {
			continue
		}
		if strings.HasPrefix(line, "#EXT-X-DISCONTINUITY-SEQUENCE") {
			continue
		}
		if strings.HasPrefix(line, "#EXT") {
			tags = append(tags, line)
			if strings.HasPrefix(line, "#EXT-X-TARGETDURATION") {
				tags = append(tags, "#EXT-X-MEDIA-SEQUENCE:0")
				tags = append(tags, "#EXT-X-DISCONTINUITY-SEQUENCE:0")
			}
			if strings.HasPrefix(line, "#EXT-X-DISCONTINUITY") {
				discontinuitySequence += 1
			}
		}
		// handle media line.
		if !strings.HasPrefix(line, "#") {

			// extract last #EXTINF tag.
			var lastInfTag string
			for i := len(tags) - 1; i >= 0; i-- {
				tag := tags[i]
				if strings.HasPrefix(tag, "#EXTINF") {
					lastInfTag = tag
					break
				}
			}

			// extract duration.
			fmt.Println("lastInfTag")
			fmt.Println(lastInfTag)
			duration, err := strconv.ParseFloat(lastInfTag[8:len(lastInfTag)-1], 64)
			if err != nil {
				return MediaPlaylist{}, nil
			}
			durationMs := duration * 1000

			// extract file extension.
			fileExtension := line[strings.LastIndex(line, "."):]

			// extract request type.
			var lastTag string
			if 0 < len(tags) {
				lastTag = tags[len(tags)-1]
			}
			var requestType RequestType
			if strings.HasPrefix(lastTag, "#EXT-X-BYTERANGE") {
				requestType = ByteRange
			} else {
				requestType = SegmentBySegment
			}

			// extract container format.
			var containerFormat ContainerFormat
			switch {
			case strings.HasSuffix(line, TsFileExtension):
				containerFormat = Ts
			case strings.HasSuffix(line, Mp4FileExtension):
				containerFormat = Fmp4
			case strings.HasPrefix(line[len(line)-4:], M4FileExtensionPrefix):
				containerFormat = Fmp4
			case strings.HasPrefix(line[len(line)-5:], Mp4FileExtensionPrefix):
				containerFormat = Fmp4
			case strings.HasPrefix(line[len(line)-5:], CmfFileExtensionPrefix):
				containerFormat = Fmp4
			default:
				return MediaPlaylist{}, nil
			}

			segment := Segment{
				Path:                  line,
				DurationMs:            durationMs,
				DiscontinuitySequence: discontinuitySequence,
				FileExtension:         fileExtension,
				ContainerFormat:       containerFormat,
				RequestType:           requestType,
			}
			segments = append(segments, segment)
			totalDurationMs += durationMs
		}
	}
	mediaPlaylist := MediaPlaylist{
		Path:                    path,
		Tags:                    tags,
		Segments:                segments,
		TotalDurationMs:         totalDurationMs,
		TotalDiscontinuityCount: discontinuitySequence,
	}
	return mediaPlaylist, nil
}
