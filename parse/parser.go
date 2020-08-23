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
		// append EXT-X-MEDIA-SEQUENCE, EXT-X-DISCONTINUITY-SEQUENCE while ignoring pre-existed those tags.
		if strings.HasPrefix(line, "#EXT-X-MEDIA-SEQUENCE") {
			continue
		}
		if strings.HasPrefix(line, "#EXT-X-DISCONTINUITY-SEQUENCE") {
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
	for {
		readBytes, _, err := reader.ReadLine()
		line := string(readBytes)
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
		if strings.HasPrefix(line, "#EXT") {
			tags = append(tags, line)
			if strings.HasPrefix(line, "#EXT-X-TARGETDURATION") {
				tags = append(tags, "#EXT-X-MEDIA-SEQUENCE:0")
				tags = append(tags, "#EXT-X-DISCONTINUITY-SEQUENCE:0")
			}
		}
		if !strings.HasPrefix(line, "#") {
			for i := len(tags) - 1; i >= 0; i-- {
				tag := tags[i]
				if strings.HasPrefix(tag, "#EXT") {
					duration, err := strconv.ParseFloat(tag[8:len(tag)-1], 64)
					if err != nil {
						fmt.Println(err)
						return MediaPlaylist{}, nil
					}
					durationMs := duration * 1000
					segment := Segment{Path: line, DurationMs: durationMs}
					segments = append(segments, segment)
					totalDurationMs += durationMs
					break
				}
			}
		}
	}
	mediaPlaylist := MediaPlaylist{
		Path:            path,
		Tags:            tags,
		Segments:        segments,
		TotalDurationMs: totalDurationMs,
	}
	return mediaPlaylist, nil
}
