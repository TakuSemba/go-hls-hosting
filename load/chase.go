package load

import (
	"github.com/TakuSemba/go-media-hosting/parse"
	"strconv"
)

type ChaseLoader struct {
	DefaultLoader
	MasterPlaylist parse.MasterPlaylist
}

func (v *ChaseLoader) LoadMediaPlaylist(index int) ([]byte, error) {
	return []byte("ChaseLoader: LoadMediaPlaylist " + strconv.Itoa(index)), nil
}
