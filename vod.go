package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

type VodLoader struct {
	MasterPlaylist MasterPlaylist
}

func (v *VodLoader) loadMasterPlaylist() ([]byte, error) {
	var masterPlaylist []byte
	var mediaPlaylistCount = 0
	for _, tag := range v.MasterPlaylist.Tags {
		masterPlaylist = append(masterPlaylist, tag...)
		masterPlaylist = append(masterPlaylist, '\n')
		if strings.HasPrefix(tag, "#EXT-X-STREAM-INF") {
			masterPlaylist = append(masterPlaylist, strconv.Itoa(mediaPlaylistCount)+"/playlist.m3u8"...)
			masterPlaylist = append(masterPlaylist, '\n')
			mediaPlaylistCount += 1
		}
	}
	return masterPlaylist, nil
}

func (v *VodLoader) loadMediaPlaylist(index int) ([]byte, error) {
	var mediaPlaylist []byte
	var tsCount = 0
	for _, tag := range v.MasterPlaylist.MediaPlaylists[index].Tags {
		mediaPlaylist = append(mediaPlaylist, tag...)
		mediaPlaylist = append(mediaPlaylist, '\n')
		if strings.HasPrefix(tag, "#EXTINF") {
			mediaPlaylist = append(mediaPlaylist, strconv.Itoa(tsCount)+".ts"...)
			mediaPlaylist = append(mediaPlaylist, '\n')
			tsCount += 1
		}
	}
	return mediaPlaylist, nil
}

func (v *VodLoader) loadSegment(mediaPlaylistIndex, segmentIndex int) ([]byte, error) {
	mediaPlaylistPath := v.MasterPlaylist.MediaPlaylists[mediaPlaylistIndex].Path
	segmentPath := v.MasterPlaylist.MediaPlaylists[mediaPlaylistIndex].Segments[segmentIndex].Path
	fmt.Println(filepath.Join(filepath.Dir(mediaPlaylistPath), segmentPath))
	segment, err := ioutil.ReadFile(filepath.Join(filepath.Dir(mediaPlaylistPath), segmentPath))
	if err != nil {
		return []byte{}, nil
	}
	return segment, nil
}
