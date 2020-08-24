package load

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLoadTsLiveMediaPlaylist(t *testing.T) {
	loader := NewLiveLoader(FakeTsMasterPlayList)
	loader.StartedAt = time.Now().Add(time.Duration(-3 * 1000 * 1000 * 1000))
	loader.WindowDurationMs = 9 * 1000
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := `
		#EXTM3U
		#EXT-X-VERSION:4
		#EXT-X-PLAYLIST-TYPE:EVENT
		#EXT-X-INDEPENDENT-SEGMENTS
		#EXT-X-TARGETDURATION:5
		#EXT-X-MEDIA-SEQUENCE:1
		#EXT-X-DISCONTINUITY-SEQUENCE:0
		#EXTINF:4,
		1.ts
		#EXTINF:5,
		2.ts
	`
	if assert.NoError(t, err) {
		assert.Equal(t, trimIndent(t, mediaPlaylist), string(actual))
	}
}

func TestLoadFmp4LiveMediaPlaylist(t *testing.T) {
	loader := NewLiveLoader(FakeFmp4MasterPlayList)
	loader.StartedAt = time.Now().Add(time.Duration(-3 * 1000 * 1000 * 1000))
	loader.WindowDurationMs = 9 * 1000
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := `
		#EXTM3U
		#EXT-X-VERSION:4
		#EXT-X-PLAYLIST-TYPE:EVENT
		#EXT-X-INDEPENDENT-SEGMENTS
		#EXT-X-TARGETDURATION:5
		#EXT-X-MEDIA-SEQUENCE:1
		#EXT-X-DISCONTINUITY-SEQUENCE:0
		#EXTINF:4,
		1.mp4
		#EXTINF:5,
		2.mp4
	`
	if assert.NoError(t, err) {
		assert.Equal(t, trimIndent(t, mediaPlaylist), string(actual))
	}
}

func TestLoadByteRangeLiveMediaPlaylist(t *testing.T) {
	loader := NewLiveLoader(FakeByteRangeMasterPlayList)
	loader.StartedAt = time.Now().Add(time.Duration(-3 * 1000 * 1000 * 1000))
	loader.WindowDurationMs = 9 * 1000
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := `
		#EXTM3U
		#EXT-X-VERSION:4
		#EXT-X-PLAYLIST-TYPE:EVENT
		#EXT-X-INDEPENDENT-SEGMENTS
		#EXT-X-TARGETDURATION:5
		#EXT-X-MEDIA-SEQUENCE:1
		#EXT-X-DISCONTINUITY-SEQUENCE:0
		#EXTINF:4,
		#EXT-X-BYTERANGE:400000@300100
		0.mp4
		#EXTINF:5,
		#EXT-X-BYTERANGE:500000@700100
		0.mp4
	`
	if assert.NoError(t, err) {
		assert.Equal(t, trimIndent(t, mediaPlaylist), string(actual))
	}
}

func TestLoadLiveMediaPlaylistWhenExceeded(t *testing.T) {
	loader := NewLiveLoader(FakeTsMasterPlayList)
	loader.StartedAt = time.Now().Add(time.Duration(-7 * 1000 * 1000 * 1000))
	loader.WindowDurationMs = 8 * 1000
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := `
		#EXTM3U
		#EXT-X-VERSION:4
		#EXT-X-PLAYLIST-TYPE:EVENT
		#EXT-X-INDEPENDENT-SEGMENTS
		#EXT-X-TARGETDURATION:5
		#EXT-X-MEDIA-SEQUENCE:2
		#EXT-X-DISCONTINUITY-SEQUENCE:0
		#EXTINF:5,
		2.ts
		#EXT-X-DISCONTINUITY
		#EXTINF:3,
		0.ts
	`
	if assert.NoError(t, err) {
		assert.Equal(t, trimIndent(t, mediaPlaylist), string(actual))
	}
}

func TestLoadLiveMediaPlaylistWhenLooped(t *testing.T) {
	loader := NewLiveLoader(FakeTsMasterPlayList)
	loader.StartedAt = time.Now().Add(time.Duration(-12 * 1000 * 1000 * 1000))
	loader.WindowDurationMs = 7 * 1000
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := `
		#EXTM3U
		#EXT-X-VERSION:4
		#EXT-X-PLAYLIST-TYPE:EVENT
		#EXT-X-INDEPENDENT-SEGMENTS
		#EXT-X-TARGETDURATION:5
		#EXT-X-MEDIA-SEQUENCE:3
		#EXT-X-DISCONTINUITY-SEQUENCE:1
		#EXTINF:3,
		0.ts
		#EXTINF:4,
		1.ts
	`
	if assert.NoError(t, err) {
		assert.Equal(t, trimIndent(t, mediaPlaylist), string(actual))
	}
}
