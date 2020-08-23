package load

import (
	"github.com/TakuSemba/go-media-hosting/media"
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
	original := v.MasterPlaylist.MediaPlaylists[index]
	segment := original.Segments[segmentIndex]

	for _, tag := range v.MasterPlaylist.MediaPlaylists[index].Tags {
		switch {
		// append #EXT-X-PLAYLIST-TYPE:EVENT
		case strings.HasPrefix(tag, "#EXT-X-PLAYLIST-TYPE"):
			mediaPlaylist = append(mediaPlaylist, "#EXT-X-PLAYLIST-TYPE:EVENT"...)
			mediaPlaylist = append(mediaPlaylist, '\n')

		// append #EXTINF / #EXT-X-BYTERANGE
		case strings.HasPrefix(tag, media.TagMediaDuration) || strings.HasPrefix(tag, media.TagByteRange):
			mediaPlaylist = append(mediaPlaylist, tag...)
			mediaPlaylist = append(mediaPlaylist, '\n')

			switch segment.RequestType {
			// append media line for segment
			case parse.SegmentBySegment:
				if strings.HasPrefix(tag, media.TagMediaDuration) && aggregatedTimeMs < elapsedTimeMs {
					mediaPlaylist = append(mediaPlaylist, tag...)
					mediaPlaylist = append(mediaPlaylist, '\n')
					mediaPlaylist = append(mediaPlaylist, strconv.Itoa(segmentIndex)+".ts"...)
					mediaPlaylist = append(mediaPlaylist, '\n')
					aggregatedTimeMs += segment.DurationMs
					segmentIndex += 1
				}
			// append media line for byte-range
			case parse.ByteRange:
				if strings.HasPrefix(tag, media.TagByteRange) && aggregatedTimeMs < elapsedTimeMs {
					mediaPlaylist = append(mediaPlaylist, strconv.Itoa(segment.DiscontinuitySequence)+segment.FileExtension...)
					mediaPlaylist = append(mediaPlaylist, '\n')
					aggregatedTimeMs += segment.DurationMs
				}
			}

		// ignore #EXT-X-ENDLIST if needed
		case strings.HasPrefix(tag, "#EXT-X-ENDLIST"):
			if segmentIndex == len(original.Segments)-1 {
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
