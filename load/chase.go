package load

import (
	"github.com/TakuSemba/go-media-hosting/parse"
	"strconv"
	"strings"
	"time"
)

type ChaseLoader struct {
	DefaultLoader
	MasterPlaylist parse.MasterPlaylist
	StartedAt      time.Time
}

func NewChaseLoader(original parse.MasterPlaylist) ChaseLoader {
	return ChaseLoader{
		DefaultLoader:  NewDefaultLoader(original),
		MasterPlaylist: original,
		StartedAt:      time.Now(),
	}
}

func (v *ChaseLoader) LoadMediaPlaylist(index int) ([]byte, error) {
	var mediaPlaylist []byte
	var segmentIndex = 0
	aggregatedTimeMs := float64(0)
	elapsedTimeMs := float64(time.Now().Sub(v.StartedAt).Milliseconds())
	for _, tag := range v.MasterPlaylist.MediaPlaylists[index].Tags {
		switch {
		case strings.HasPrefix(tag, "#EXT-X-PLAYLIST-TYPE"):
			mediaPlaylist = append(mediaPlaylist, "#EXT-X-PLAYLIST-TYPE:EVENT"...)
			mediaPlaylist = append(mediaPlaylist, '\n')
		case strings.HasPrefix(tag, "#EXTINF"):
			if aggregatedTimeMs < elapsedTimeMs {
				mediaPlaylist = append(mediaPlaylist, tag...)
				mediaPlaylist = append(mediaPlaylist, '\n')
				mediaPlaylist = append(mediaPlaylist, strconv.Itoa(segmentIndex)+".ts"...)
				mediaPlaylist = append(mediaPlaylist, '\n')
				aggregatedTimeMs += v.MasterPlaylist.MediaPlaylists[index].Segments[segmentIndex].DurationMs
				segmentIndex += 1
			}
		case strings.HasPrefix(tag, "#EXT-X-ENDLIST"):
			if segmentIndex == len(v.MasterPlaylist.MediaPlaylists[index].Segments)-1 {
				mediaPlaylist = append(mediaPlaylist, tag...)
				mediaPlaylist = append(mediaPlaylist, '\n')
			}
		default:
			mediaPlaylist = append(mediaPlaylist, tag...)
			mediaPlaylist = append(mediaPlaylist, '\n')
		}
	}
	return mediaPlaylist, nil
}
