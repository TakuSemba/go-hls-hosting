package load

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLoadTsChaseMediaPlaylist(t *testing.T) {
	loader := NewChaseLoader(FakeTsMasterPlayList)
	loader.StartedAt = time.Now().Add(time.Duration(-4 * 1000 * 1000 * 1000))
	loader.InitialWindowDurationMs = 3 * 1000
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := `
		#EXTM3U
		#EXT-X-VERSION:4
		#EXT-X-PLAYLIST-TYPE:EVENT
		#EXT-X-INDEPENDENT-SEGMENTS
		#EXT-X-TARGETDURATION:5
		#EXT-X-MEDIA-SEQUENCE:0
		#EXT-X-DISCONTINUITY-SEQUENCE:0
		#EXTINF:3,
		0.ts
		#EXTINF:4,
		1.ts
	`
	if assert.NoError(t, err) {
		assert.Equal(t, trimIndent(t, mediaPlaylist), string(actual))
	}
}

func TestLoadFmp4ChaseMediaPlaylist(t *testing.T) {
	loader := NewChaseLoader(FakeFmp4MasterPlayList)
	loader.StartedAt = time.Now().Add(time.Duration(-4 * 1000 * 1000 * 1000))
	loader.InitialWindowDurationMs = 3 * 1000
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := `
		#EXTM3U
		#EXT-X-VERSION:4
		#EXT-X-PLAYLIST-TYPE:EVENT
		#EXT-X-INDEPENDENT-SEGMENTS
		#EXT-X-TARGETDURATION:5
		#EXT-X-MEDIA-SEQUENCE:0
		#EXT-X-DISCONTINUITY-SEQUENCE:0
		#EXTINF:3,
		0.mp4
		#EXTINF:4,
		1.mp4
	`
	if assert.NoError(t, err) {
		assert.Equal(t, trimIndent(t, mediaPlaylist), string(actual))
	}
}

func TestLoadByteRangeChaseMediaPlaylist(t *testing.T) {
	loader := NewChaseLoader(FakeByteRangeMasterPlayList)
	loader.StartedAt = time.Now().Add(time.Duration(-4 * 1000 * 1000 * 1000))
	loader.InitialWindowDurationMs = 3 * 1000
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := `
		#EXTM3U
		#EXT-X-VERSION:4
		#EXT-X-PLAYLIST-TYPE:EVENT
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
	`
	if assert.NoError(t, err) {
		assert.Equal(t, trimIndent(t, mediaPlaylist), string(actual))
	}
}

func TestLoadChaseMediaPlaylistWhenEnded(t *testing.T) {
	loader := NewChaseLoader(FakeTsMasterPlayList)
	loader.StartedAt = time.Now()
	loader.InitialWindowDurationMs = 12 * 1000
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := `
		#EXTM3U
		#EXT-X-VERSION:4
		#EXT-X-PLAYLIST-TYPE:EVENT
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
