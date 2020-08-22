package load

import (
	"github.com/TakuSemba/go-media-hosting/parse"
	"strconv"
)

type LiveLoader struct {
	MasterPlaylist parse.MasterPlaylist
}

func (v *LiveLoader) LoadMasterPlaylist() ([]byte, error) {
	return []byte("LiveLoader: LoadMasterPlaylist"), nil
}

func (v *LiveLoader) LoadMediaPlaylist(index int) ([]byte, error) {
	return []byte("LiveLoader: LoadMediaPlaylist " + strconv.Itoa(index)), nil
}

func (v *LiveLoader) LoadSegment(mediaPlaylistIndex, segmentIndex int) ([]byte, error) {
	return []byte("LiveLoader: LoadSegment " + strconv.Itoa(segmentIndex)), nil
}
