package load

import (
	"github.com/TakuSemba/go-media-hosting/parse"
	"strconv"
	"time"
)

type LiveLoader struct {
	DefaultLoader
	MasterPlaylist parse.MasterPlaylist
	StartedAt      time.Time
}

func NewLiveLoader(original parse.MasterPlaylist) LiveLoader {
	return LiveLoader{
		MasterPlaylist: original,
		StartedAt:      time.Now(),
	}
}

func (v *LiveLoader) LoadMediaPlaylist(index int) ([]byte, error) {
	return []byte("LiveLoader: LoadMediaPlaylist " + strconv.Itoa(index)), nil
}
