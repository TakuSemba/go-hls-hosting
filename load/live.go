package load

import (
	"github.com/TakuSemba/go-media-hosting/parse"
	"strconv"
	"strings"
	"time"
)

type LiveLoader struct {
	DefaultLoader
	MasterPlaylist   parse.MasterPlaylist
	StartedAt        time.Time
	WindowDurationMs float64
}

func NewLiveLoader(original parse.MasterPlaylist) LiveLoader {
	return LiveLoader{
		DefaultLoader:    NewDefaultLoader(original),
		MasterPlaylist:   original,
		StartedAt:        time.Now(),
		WindowDurationMs: 20 * 1000,
	}
}

func (v *LiveLoader) LoadMediaPlaylist(index int) ([]byte, error) {
	var mediaPlaylist []byte
	totalDurationMs := v.MasterPlaylist.MediaPlaylists[index].TotalDurationMs
	elapsedTimeMs := time.Now().UnixNano()/1e6 - v.StartedAt.UnixNano()/1e6
	adjustedElapsedTimeMs := float64(uint64(elapsedTimeMs) % uint64(totalDurationMs))

	segmentIndex := -1
	aggregatedTimeMs := float64(0)
	for aggregatedTimeMs < adjustedElapsedTimeMs {
		segmentIndex += 1
		aggregatedTimeMs += v.MasterPlaylist.MediaPlaylists[index].Segments[segmentIndex].DurationMs
	}

	aggregatedTimeMs = float64(0)
	for _, tag := range v.MasterPlaylist.MediaPlaylists[index].Tags {
		if strings.HasPrefix(tag, "#EXTINF") {
			if aggregatedTimeMs < v.WindowDurationMs {
				mediaPlaylist = append(mediaPlaylist, tag...)
				mediaPlaylist = append(mediaPlaylist, '\n')
				mediaPlaylist = append(mediaPlaylist, strconv.Itoa(segmentIndex)+".ts"...)
				mediaPlaylist = append(mediaPlaylist, '\n')
				aggregatedTimeMs += v.MasterPlaylist.MediaPlaylists[index].Segments[segmentIndex].DurationMs
				segmentIndex += 1
			}
		} else if strings.HasPrefix(tag, "#EXT-X-ENDLIST") {
			continue
		} else {
			mediaPlaylist = append(mediaPlaylist, tag...)
			mediaPlaylist = append(mediaPlaylist, '\n')
		}
	}
	return mediaPlaylist, nil
}
