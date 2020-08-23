package load

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLoadTsChaseMediaPlaylist(t *testing.T) {
	loader := NewChaseLoader(FakeTsMasterPlayList)
	// 16 seconds before
	loader.StartedAt = time.Now().Add(time.Duration(-16 * 1000 * 1000 * 1000))
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := "#EXTM3U\n" +
		"#EXT-X-VERSION:4\n" +
		"#EXT-X-PLAYLIST-TYPE:EVENT\n" +
		"#EXT-X-INDEPENDENT-SEGMENTS\n" +
		"#EXT-X-TARGETDURATION:5\n" +
		"#EXT-X-MEDIA-SEQUENCE:0\n" +
		"#EXT-X-DISCONTINUITY-SEQUENCE:0\n" +
		"#EXTINF:3,\n" +
		"0.ts\n" +
		"#EXTINF:4,\n" +
		"1.ts\n" +
		"#EXTINF:5,\n" +
		"2.ts\n" +
		"#EXTINF:3,\n" +
		"3.ts\n"
	if assert.NoError(t, err) {
		assert.Equal(t, mediaPlaylist, string(actual))
	}
}

func TestLoadFmp4ChaseMediaPlaylist(t *testing.T) {
	loader := NewChaseLoader(FakeFmp4MasterPlayList)
	// 16 seconds before
	loader.StartedAt = time.Now().Add(time.Duration(-16 * 1000 * 1000 * 1000))
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := "#EXTM3U\n" +
		"#EXT-X-VERSION:4\n" +
		"#EXT-X-PLAYLIST-TYPE:EVENT\n" +
		"#EXT-X-INDEPENDENT-SEGMENTS\n" +
		"#EXT-X-TARGETDURATION:5\n" +
		"#EXT-X-MEDIA-SEQUENCE:0\n" +
		"#EXT-X-DISCONTINUITY-SEQUENCE:0\n" +
		"#EXTINF:3,\n" +
		"0.mp4\n" +
		"#EXTINF:4,\n" +
		"1.mp4\n" +
		"#EXTINF:5,\n" +
		"2.mp4\n" +
		"#EXTINF:3,\n" +
		"3.mp4\n"
	if assert.NoError(t, err) {
		assert.Equal(t, mediaPlaylist, string(actual))
	}
}

func TestLoadByteRangeChaseMediaPlaylist(t *testing.T) {
	loader := NewChaseLoader(FakeByteRangeMasterPlayList)
	// 16 seconds before
	loader.StartedAt = time.Now().Add(time.Duration(-16 * 1000 * 1000 * 1000))
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := "#EXTM3U\n" +
		"#EXT-X-VERSION:4\n" +
		"#EXT-X-PLAYLIST-TYPE:EVENT\n" +
		"#EXT-X-INDEPENDENT-SEGMENTS\n" +
		"#EXT-X-TARGETDURATION:5\n" +
		"#EXT-X-MEDIA-SEQUENCE:0\n" +
		"#EXT-X-DISCONTINUITY-SEQUENCE:0\n" +
		"#EXTINF:3,\n" +
		"#EXT-X-BYTERANGE:300000@100\n" +
		"0.mp4\n" +
		"#EXTINF:4,\n" +
		"#EXT-X-BYTERANGE:400000@300100\n" +
		"0.mp4\n" +
		"#EXTINF:5,\n" +
		"#EXT-X-BYTERANGE:500000@700100\n" +
		"0.mp4\n" +
		"#EXTINF:3,\n" +
		"#EXT-X-BYTERANGE:300000@1200100\n" +
		"0.mp4\n"

	if assert.NoError(t, err) {
		assert.Equal(t, mediaPlaylist, string(actual))
	}
}
