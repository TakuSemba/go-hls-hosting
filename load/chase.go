package load

import (
	"github.com/TakuSemba/go-media-hosting/parse"
	"strconv"
	"time"
)

type ChaseLoader struct {
	DefaultLoader
	MasterPlaylist parse.MasterPlaylist
	StartedAt      time.Time
}

func NewChaseLoader(original parse.MasterPlaylist) ChaseLoader {
	return ChaseLoader{
		MasterPlaylist: original,
		StartedAt:      time.Now(),
	}
}

func (v *ChaseLoader) LoadMediaPlaylist(index int) ([]byte, error) {
	return []byte("ChaseLoader: LoadMediaPlaylist " + strconv.Itoa(index)), nil
}
