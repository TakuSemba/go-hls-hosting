package load

import (
	"github.com/TakuSemba/go-hls-hosting/media"
	"github.com/TakuSemba/go-hls-hosting/parse"
	"strconv"
	"strings"
	"time"
)

type ChaseLoader struct {
	DefaultLoader
	MasterPlaylist          parse.MasterPlaylist
	StartedAt               time.Time
	InitialWindowDurationMs float64
}

func NewChaseLoader(original parse.MasterPlaylist) ChaseLoader {
	return ChaseLoader{
		DefaultLoader:           NewDefaultLoader(original),
		MasterPlaylist:          original,
		StartedAt:               time.Now(),
		InitialWindowDurationMs: 40 * 1000,
	}
}

func (v *ChaseLoader) LoadMediaPlaylist(index int) ([]byte, error) {
	var mediaPlaylist []byte
	var segmentIndex = 0
	aggregatedTimeMs := float64(0)
	elapsedTimeMs := float64(time.Now().Sub(v.StartedAt).Milliseconds())
	original := v.MasterPlaylist.MediaPlaylists[index]

	windowDurationMs := elapsedTimeMs + v.InitialWindowDurationMs

	for _, tag := range v.MasterPlaylist.MediaPlaylists[index].Tags {
		switch {
		// append #EXT-X-PLAYLIST-TYPE:EVENT.
		case strings.HasPrefix(tag, "#EXT-X-PLAYLIST-TYPE"):
			mediaPlaylist = append(mediaPlaylist, "#EXT-X-PLAYLIST-TYPE:EVENT"...)
			mediaPlaylist = append(mediaPlaylist, '\n')

		case strings.HasPrefix(tag, media.TagDiscontinuity):
			// ignore if next segment is out of window.
			if windowDurationMs < aggregatedTimeMs+original.Segments[segmentIndex].DurationMs {
				continue
			}
			mediaPlaylist = append(mediaPlaylist, tag...)
			mediaPlaylist = append(mediaPlaylist, '\n')

		// append #EXTINF / #EXT-X-BYTERANGE.
		case strings.HasPrefix(tag, media.TagMediaDuration) || strings.HasPrefix(tag, media.TagByteRange):
			segment := original.Segments[segmentIndex]

			if aggregatedTimeMs+segment.DurationMs < windowDurationMs {
				mediaPlaylist = append(mediaPlaylist, tag...)
				mediaPlaylist = append(mediaPlaylist, '\n')

				switch segment.RequestType {
				// append media line for segment.
				case parse.SegmentBySegment:
					if strings.HasPrefix(tag, media.TagMediaDuration) {
						mediaPlaylist = append(mediaPlaylist, strconv.Itoa(segmentIndex)+segment.FileExtension...)
						mediaPlaylist = append(mediaPlaylist, '\n')
						aggregatedTimeMs += segment.DurationMs
						segmentIndex += 1
					}
				// append media line for byte-range.
				case parse.ByteRange:
					if strings.HasPrefix(tag, media.TagByteRange) {
						mediaPlaylist = append(mediaPlaylist, strconv.Itoa(segment.DiscontinuitySequence)+segment.FileExtension...)
						mediaPlaylist = append(mediaPlaylist, '\n')
						aggregatedTimeMs += segment.DurationMs
						segmentIndex += 1
					}
				}
			}

		// ignore #EXT-X-ENDLIST if needed.
		case strings.HasPrefix(tag, "#EXT-X-ENDLIST"):
			if len(original.Segments)-1 < segmentIndex {
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
