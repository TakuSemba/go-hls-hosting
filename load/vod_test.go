package load

import (
	"testing"
)

func TestLoadMediaPlaylist(t *testing.T) {
	original := FakeMasterPlayList
	loader := NewVodLoader(original)
	actual, err := loader.LoadMediaPlaylist(0)
	if err != nil {
		t.Fatalf("error happened: %#v", err)
	}
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
	if mediaPlaylist != string(actual) {
		t.Errorf("exspected: %v, actual: %v", mediaPlaylist, string(actual))
	}
}
