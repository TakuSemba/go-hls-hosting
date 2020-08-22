package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strings"
)

const TagPrefix = "#EXT"

type Parser struct{}

func (p *Parser) Parse(path string) (MasterPlaylist, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return MasterPlaylist{}, err
	}
	return p.ParseMasterPlaylist(data)
}

func (p *Parser) ParseMasterPlaylist(data []byte) (MasterPlaylist, error) {
	reader := bufio.NewReader(bytes.NewReader(data))
	var tags []string
	var mediaPlaylists []MediaPlaylist
	for {
		readBytes, _, err := reader.ReadLine()
		line := string(readBytes)
		if len(line) == 0 {
			break
		}
		if strings.HasPrefix(line, TagPrefix) {
			tags = append(tags, line)
		}
		if !strings.HasPrefix(line, "#") {
			data, err := ioutil.ReadFile(line)
			if err != nil {
				break
			}
			mediaPlaylist, err := p.ParseMediaPlaylist(data)
			if err != nil {
				return MasterPlaylist{}, err
			}
			mediaPlaylists = append(mediaPlaylists, mediaPlaylist)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return MasterPlaylist{}, err
		}
	}
	masterPlaylist := MasterPlaylist{
		Tags:           tags,
		MediaPlaylists: mediaPlaylists,
	}
	return masterPlaylist, nil
}

func (p *Parser) ParseMediaPlaylist(data []byte) (MediaPlaylist, error) {
	reader := bufio.NewReader(bytes.NewReader(data))
	var tags []string
	var segments []Segment
	for {
		readBytes, _, err := reader.ReadLine()
		line := string(readBytes)
		if len(line) == 0 {
			break
		}
		if strings.HasPrefix(line, TagPrefix) {
			tags = append(tags, line)
		}
		if !strings.HasPrefix(line, "#") {
			segment := Segment{Uri: line}
			segments = append(segments, segment)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return MediaPlaylist{}, err
		}
	}
	mediaPlaylist := MediaPlaylist{
		Tags:     tags,
		Segments: segments,
	}
	return mediaPlaylist, nil
}
