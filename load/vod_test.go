package load

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadTsVodMediaPlaylist(t *testing.T) {
	loader := NewVodLoader(FakeTsMasterPlayList)
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := `
		#EXTM3U
		#EXT-X-VERSION:4
		#EXT-X-PLAYLIST-TYPE:VOD
		#EXT-X-INDEPENDENT-SEGMENTS
		#EXT-X-TARGETDURATION:5
		#EXT-X-MEDIA-SEQUENCE:0
		#EXT-X-DISCONTINUITY-SEQUENCE:0
		#EXTINF:3,
		0.ts
		#EXTINF:4,
		1.ts
		#EXTINF:5,
		2.ts
		#EXT-X-ENDLIST
	`

	if assert.NoError(t, err) {
		assert.Equal(t, trimIndent(t, mediaPlaylist), string(actual))
	}
}

func TestLoadFmp4VodMediaPlaylist(t *testing.T) {
	loader := NewVodLoader(FakeFmp4MasterPlayList)
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := `
		#EXTM3U
		#EXT-X-VERSION:4
		#EXT-X-PLAYLIST-TYPE:VOD
		#EXT-X-INDEPENDENT-SEGMENTS
		#EXT-X-TARGETDURATION:5
		#EXT-X-MEDIA-SEQUENCE:0
		#EXT-X-DISCONTINUITY-SEQUENCE:0
		#EXTINF:3,
		0.mp4
		#EXTINF:4,
		1.mp4
		#EXTINF:5,
		2.mp4
		#EXT-X-ENDLIST
	`

	if assert.NoError(t, err) {
		assert.Equal(t, trimIndent(t, mediaPlaylist), string(actual))
	}
}

func TestLoadByteRangeVodMediaPlaylist(t *testing.T) {
	loader := NewVodLoader(FakeByteRangeMasterPlayList)
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := `
		#EXTM3U
		#EXT-X-VERSION:4
		#EXT-X-PLAYLIST-TYPE:VOD
		#EXT-X-INDEPENDENT-SEGMENTS
		#EXT-X-TARGETDURATION:5
		#EXT-X-MEDIA-SEQUENCE:0
		#EXT-X-DISCONTINUITY-SEQUENCE:0
		#EXTINF:3,
		#EXT-X-BYTERANGE:300000@100
		0.mp4
		#EXTINF:4,
		#EXT-X-BYTERANGE:400000@300100
		0.mp4
		#EXTINF:5,
		#EXT-X-BYTERANGE:500000@700100
		0.mp4
		#EXT-X-ENDLIST
	`

	if assert.NoError(t, err) {
		assert.Equal(t, trimIndent(t, mediaPlaylist), string(actual))
	}
}
