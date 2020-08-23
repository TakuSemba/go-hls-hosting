package load

import (
	"testing"
	"time"
)

func TestLiveLoadMediaPlaylist(t *testing.T) {
	original := FakeMasterPlayList
	loader := NewLiveLoader(original)
	// 5 minutes before
	loader.StartedAt = time.Now().Add(time.Duration(-5 * 60 * 1000 * 1000 * 1000))
	actual, err := loader.LoadMediaPlaylist(0)
	if err != nil {
		t.Fatalf("error happened: %#v", err)
	}
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
		"7.ts\n" +
		"#EXTINF:5,\n" +
		"8.ts\n"
	if mediaPlaylist != string(actual) {
		t.Errorf("exspected: %v, actual: %v", mediaPlaylist, string(actual))
	}
}
