package load

import (
	"github.com/TakuSemba/go-media-hosting/parse"
	"strconv"
)

type ChaseLoader struct {
	MasterPlaylist parse.MasterPlaylist
}

func (v *ChaseLoader) LoadMasterPlaylist() ([]byte, error) {
	return []byte("ChaseLoader: LoadMasterPlaylist"), nil
}

func (v *ChaseLoader) LoadMediaPlaylist(index int) ([]byte, error) {
	return []byte("ChaseLoader: LoadMediaPlaylist " + strconv.Itoa(index)), nil
}

func (v *ChaseLoader) LoadSegment(mediaPlaylistIndex, segmentIndex int) ([]byte, error) {
	return []byte("ChaseLoader: LoadSegment " + strconv.Itoa(segmentIndex)), nil
}
