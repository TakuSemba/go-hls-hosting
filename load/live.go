package load

import (
	"github.com/TakuSemba/go-media-hosting/parse"
	"strconv"
)

type LiveLoader struct {
	DefaultLoader
	MasterPlaylist parse.MasterPlaylist
}

func (v *LiveLoader) LoadMediaPlaylist(index int) ([]byte, error) {
	return []byte("LiveLoader: LoadMediaPlaylist " + strconv.Itoa(index)), nil
}
