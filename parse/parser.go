package parse

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"path/filepath"
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
		}
		if !strings.HasPrefix(line, "#") {
			segment := Segment{Path: line}
			segments = append(segments, segment)
		}
	}
	mediaPlaylist := MediaPlaylist{
		Path:     path,
		Tags:     tags,
		Segments: segments,
	}
	return mediaPlaylist, nil
}
