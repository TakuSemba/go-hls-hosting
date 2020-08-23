package load

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLoadTsLiveMediaPlaylist(t *testing.T) {
	loader := NewLiveLoader(FakeTsMasterPlayList)
	// 5 minutes before
	loader.StartedAt = time.Now().Add(time.Duration(-5 * 60 * 1000 * 1000 * 1000))
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := "#EXTM3U\n" +
		"#EXT-X-VERSION:4\n" +
		"#EXT-X-PLAYLIST-TYPE:EVENT\n" +
		"#EXT-X-INDEPENDENT-SEGMENTS\n" +
		"#EXT-X-TARGETDURATION:5\n" +
		"#EXT-X-MEDIA-SEQUENCE:75\n" +
		"#EXT-X-DISCONTINUITY-SEQUENCE:0\n" +
		"#EXTINF:3,\n" +
		"3.ts\n" +
		"#EXTINF:4,\n" +
		"4.ts\n" +
		"#EXTINF:5,\n" +
		"5.ts\n" +
		"#EXTINF:3,\n" +
		"6.ts\n" +
		"#EXTINF:4,\n" +
		"7.ts\n"
	if assert.NoError(t, err) {
		assert.Equal(t, mediaPlaylist, string(actual))
	}
}

func TestLoadFmp4LiveMediaPlaylist(t *testing.T) {
	loader := NewLiveLoader(FakeFmp4MasterPlayList)
	// 5 minutes before
	loader.StartedAt = time.Now().Add(time.Duration(-5 * 60 * 1000 * 1000 * 1000))
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := "#EXTM3U\n" +
		"#EXT-X-VERSION:4\n" +
		"#EXT-X-PLAYLIST-TYPE:EVENT\n" +
		"#EXT-X-INDEPENDENT-SEGMENTS\n" +
		"#EXT-X-TARGETDURATION:5\n" +
		"#EXT-X-MEDIA-SEQUENCE:75\n" +
		"#EXT-X-DISCONTINUITY-SEQUENCE:0\n" +
		"#EXTINF:3,\n" +
		"3.mp4\n" +
		"#EXTINF:4,\n" +
		"4.mp4\n" +
		"#EXTINF:5,\n" +
		"5.mp4\n" +
		"#EXTINF:3,\n" +
		"6.mp4\n" +
		"#EXTINF:4,\n" +
		"7.mp4\n"

	if assert.NoError(t, err) {
		assert.Equal(t, mediaPlaylist, string(actual))
	}
}

func TestLoadByteRangeLiveMediaPlaylist(t *testing.T) {
	loader := NewLiveLoader(FakeByteRangeMasterPlayList)
	// 5 minutes before
	loader.StartedAt = time.Now().Add(time.Duration(-5 * 60 * 1000 * 1000 * 1000))
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := "#EXTM3U\n" +
		"#EXT-X-VERSION:4\n" +
		"#EXT-X-PLAYLIST-TYPE:EVENT\n" +
		"#EXT-X-INDEPENDENT-SEGMENTS\n" +
		"#EXT-X-TARGETDURATION:5\n" +
		"#EXT-X-MEDIA-SEQUENCE:75\n" +
		"#EXT-X-DISCONTINUITY-SEQUENCE:0\n" +
		"#EXTINF:3,\n" +
		"#EXT-X-BYTERANGE:300000@1200100\n" +
		"0.mp4\n" +
		"#EXTINF:4,\n" +
		"#EXT-X-BYTERANGE:400000@1500100\n" +
		"0.mp4\n" +
		"#EXTINF:5,\n" +
		"#EXT-X-BYTERANGE:500000@1900100\n" +
		"0.mp4\n" +
		"#EXTINF:3,\n" +
		"#EXT-X-BYTERANGE:300000@2400100\n" +
		"0.mp4\n" +
		"#EXTINF:4,\n" +
		"#EXT-X-BYTERANGE:400000@2700100\n" +
		"0.mp4\n"

	if assert.NoError(t, err) {
		assert.Equal(t, mediaPlaylist, string(actual))
	}
}
