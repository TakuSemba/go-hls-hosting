package load

import (
	"github.com/TakuSemba/go-media-hosting/media"
	"github.com/TakuSemba/go-media-hosting/parse"
	"strconv"
	"strings"
)

type VodLoader struct {
	DefaultLoader
	MasterPlaylist parse.MasterPlaylist
}

func NewVodLoader(original parse.MasterPlaylist) VodLoader {
	return VodLoader{
		DefaultLoader:  NewDefaultLoader(original),
		MasterPlaylist: original,
	}
}

func (v *VodLoader) LoadMediaPlaylist(index int) ([]byte, error) {
	var mediaPlaylist []byte
	var segmentIndex = 0
	for _, tag := range v.MasterPlaylist.MediaPlaylists[index].Tags {
		switch {
		// append #EXTINF / #EXT-X-BYTERANGE
		case strings.HasPrefix(tag, media.TagMediaDuration) || strings.HasPrefix(tag, media.TagByteRange):
			mediaPlaylist = append(mediaPlaylist, tag...)
			mediaPlaylist = append(mediaPlaylist, '\n')
			segment := v.MasterPlaylist.MediaPlaylists[index].Segments[segmentIndex]
			switch segment.RequestType {
			// append media line for segment
			case parse.SegmentBySegment:
				if strings.HasPrefix(tag, media.TagMediaDuration) {
					mediaPlaylist = append(mediaPlaylist, strconv.Itoa(segmentIndex)+segment.FileExtension...)
					mediaPlaylist = append(mediaPlaylist, '\n')
					segmentIndex += 1
				}

			// append media line for byte-range
			case parse.ByteRange:
				if strings.HasPrefix(tag, media.TagByteRange) {
					mediaPlaylist = append(mediaPlaylist, strconv.Itoa(segment.DiscontinuitySequence)+segment.FileExtension...)
					mediaPlaylist = append(mediaPlaylist, '\n')
				}
			}
		default:
			mediaPlaylist = append(mediaPlaylist, tag...)
			mediaPlaylist = append(mediaPlaylist, '\n')
		}
	}
	return mediaPlaylist, nil
}
