package load

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadTsVodMediaPlaylist(t *testing.T) {
	loader := NewVodLoader(FakeTsMasterPlayList)
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := "#EXTM3U\n" +
		"#EXT-X-VERSION:4\n" +
		"#EXT-X-PLAYLIST-TYPE:VOD\n" +
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
		"3.ts\n" +
		"#EXTINF:4,\n" +
		"4.ts\n" +
		"#EXTINF:5,\n" +
		"5.ts\n" +
		"#EXTINF:3,\n" +
		"6.ts\n" +
		"#EXTINF:4,\n" +
		"7.ts\n" +
		"#EXTINF:5,\n" +
		"8.ts\n" +
		"#EXT-X-ENDLIST\n"

	if assert.NoError(t, err) {
		assert.Equal(t, mediaPlaylist, string(actual))
	}
}

func TestLoadFmp4VodMediaPlaylist(t *testing.T) {
	loader := NewVodLoader(FakeFmp4MasterPlayList)
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := "#EXTM3U\n" +
		"#EXT-X-VERSION:4\n" +
		"#EXT-X-PLAYLIST-TYPE:VOD\n" +
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
		"3.mp4\n" +
		"#EXTINF:4,\n" +
		"4.mp4\n" +
		"#EXTINF:5,\n" +
		"5.mp4\n" +
		"#EXTINF:3,\n" +
		"6.mp4\n" +
		"#EXTINF:4,\n" +
		"7.mp4\n" +
		"#EXTINF:5,\n" +
		"8.mp4\n" +
		"#EXT-X-ENDLIST\n"

	if assert.NoError(t, err) {
		assert.Equal(t, mediaPlaylist, string(actual))
	}
}

func TestLoadByteRangeVodMediaPlaylist(t *testing.T) {
	loader := NewVodLoader(FakeByteRangeMasterPlayList)
	actual, err := loader.LoadMediaPlaylist(0)
	mediaPlaylist := "#EXTM3U\n" +
		"#EXT-X-VERSION:4\n" +
		"#EXT-X-PLAYLIST-TYPE:VOD\n" +
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
		"0.mp4\n" +
		"#EXTINF:5,\n" +
		"#EXT-X-BYTERANGE:500000@3100100\n" +
		"0.mp4\n" +
		"#EXT-X-ENDLIST\n"

	if assert.NoError(t, err) {
		assert.Equal(t, mediaPlaylist, string(actual))
	}
}
