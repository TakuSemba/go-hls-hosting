package load

import (
	"github.com/TakuSemba/go-media-hosting/media"
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
	segment := original.Segments[segmentIndex]
	for _, tag := range original.Tags {
		switch {
		// append #EXT-X-PLAYLIST-TYPE:EVENT
		case strings.HasPrefix(tag, "#EXT-X-PLAYLIST-TYPE"):
			mediaPlaylist = append(mediaPlaylist, "#EXT-X-PLAYLIST-TYPE:EVENT"...)
			mediaPlaylist = append(mediaPlaylist, '\n')

		// append #EXT-X-MEDIA-SEQUENCE:xx
		case strings.HasPrefix(tag, media.TagMediaSequence):
			mediaSequence := "#EXT-X-MEDIA-SEQUENCE:" + strconv.Itoa(segmentIndex)
			mediaPlaylist = append(mediaPlaylist, mediaSequence...)
			mediaPlaylist = append(mediaPlaylist, '\n')

		// append #EXT-X-DISCONTINUITY-SEQUENCE:xx
		case strings.HasPrefix(tag, media.TagDiscontinuitySequence):
			discontinuitySequence := repeatedWindowCount*original.TotalDiscontinuityCount + segment.DiscontinuitySequence
			mediaSequence := "#EXT-X-DISCONTINUITY-SEQUENCE:" + strconv.Itoa(discontinuitySequence)
			mediaPlaylist = append(mediaPlaylist, mediaSequence...)
			mediaPlaylist = append(mediaPlaylist, '\n')

		// append #EXTINF / #EXT-X-BYTERANGE
		case strings.HasPrefix(tag, media.TagMediaDuration) || strings.HasPrefix(tag, media.TagByteRange):
			mediaPlaylist = append(mediaPlaylist, tag...)
			mediaPlaylist = append(mediaPlaylist, '\n')

			switch segment.RequestType {
			// append media line for segment
			case parse.SegmentBySegment:
				if strings.HasPrefix(tag, media.TagMediaDuration) && aggregatedTimeMs < v.WindowDurationMs {
					mediaPlaylist = append(mediaPlaylist, strconv.Itoa(segmentIndex)+segment.FileExtension...)
					mediaPlaylist = append(mediaPlaylist, '\n')
					aggregatedTimeMs += segment.DurationMs
					segmentIndex += 1
					segmentIndex = segmentIndex % len(original.Segments)
				}
			// append media line for byte-range
			case parse.ByteRange:
				if strings.HasPrefix(tag, media.TagByteRange) && aggregatedTimeMs < v.WindowDurationMs {
					mediaPlaylist = append(mediaPlaylist, strconv.Itoa(segment.DiscontinuitySequence)+segment.FileExtension...)
					mediaPlaylist = append(mediaPlaylist, '\n')
					aggregatedTimeMs += segment.DurationMs
				}
			}

		// ignore #EXT-X-ENDLIST
		case strings.HasPrefix(tag, "#EXT-X-ENDLIST"):
			continue
		default:
			mediaPlaylist = append(mediaPlaylist, tag...)
			mediaPlaylist = append(mediaPlaylist, '\n')
		}
	}
	return mediaPlaylist, nil
}
