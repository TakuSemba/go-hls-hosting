package load

import (
	"github.com/TakuSemba/go-media-hosting/parse"
	"strconv"
	"strings"
)

type VodLoader struct {
	DefaultLoader
	MasterPlaylist parse.MasterPlaylist
}

func (v *VodLoader) LoadMediaPlaylist(index int) ([]byte, error) {
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
