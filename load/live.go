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
	original := v.MasterPlaylist.MediaPlaylists[index]
	totalDurationMs := original.TotalDurationMs
	elapsedTimeMs := time.Now().Sub(v.StartedAt).Milliseconds()
	adjustedElapsedTimeMs := float64(uint64(elapsedTimeMs) % uint64(totalDurationMs))
	repeatedWindowCount := int(uint64(elapsedTimeMs) / uint64(totalDurationMs))

	segmentIndex := -1
	aggregatedTimeMs := float64(0)
	for aggregatedTimeMs < adjustedElapsedTimeMs {
		segmentIndex += 1
		aggregatedTimeMs += original.Segments[segmentIndex].DurationMs
	}

	aggregatedTimeMs = float64(0)
	for _, tag := range original.Tags {
		switch {
		case strings.HasPrefix(tag, "#EXT-X-PLAYLIST-TYPE"):
			mediaPlaylist = append(mediaPlaylist, "#EXT-X-PLAYLIST-TYPE:EVENT"...)
			mediaPlaylist = append(mediaPlaylist, '\n')
		case strings.HasPrefix(tag, "#EXT-X-MEDIA-SEQUENCE"):
			mediaSequence := "#EXT-X-MEDIA-SEQUENCE:" + strconv.Itoa(segmentIndex)
			mediaPlaylist = append(mediaPlaylist, mediaSequence...)
			mediaPlaylist = append(mediaPlaylist, '\n')
		case strings.HasPrefix(tag, "#EXT-X-DISCONTINUITY-SEQUENCE"):
			discontinuitySequence := repeatedWindowCount*original.TotalDiscontinuityCount + original.Segments[segmentIndex].DiscontinuitySequence
			mediaSequence := "#EXT-X-DISCONTINUITY-SEQUENCE:" + strconv.Itoa(discontinuitySequence)
			mediaPlaylist = append(mediaPlaylist, mediaSequence...)
			mediaPlaylist = append(mediaPlaylist, '\n')
		case strings.HasPrefix(tag, "#EXTINF"):
			if aggregatedTimeMs < v.WindowDurationMs {
				mediaPlaylist = append(mediaPlaylist, tag...)
				mediaPlaylist = append(mediaPlaylist, '\n')
				mediaPlaylist = append(mediaPlaylist, strconv.Itoa(segmentIndex)+original.Segments[segmentIndex].FileExtension...)
				mediaPlaylist = append(mediaPlaylist, '\n')
				aggregatedTimeMs += original.Segments[segmentIndex].DurationMs
				segmentIndex += 1
				segmentIndex = segmentIndex % len(original.Segments)
			}
		case strings.HasPrefix(tag, "#EXT-X-ENDLIST"):
			continue
		default:
			mediaPlaylist = append(mediaPlaylist, tag...)
			mediaPlaylist = append(mediaPlaylist, '\n')
		}
	}
	return mediaPlaylist, nil
}
