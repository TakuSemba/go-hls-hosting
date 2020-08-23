package load

import (
	"github.com/TakuSemba/go-media-hosting/parse"
	"testing"
)

func TestVodLoadMasterPlaylist(t *testing.T) {
	original := parse.MasterPlaylist{
		Path: "testMasterPlaylist.m3u8",
		Tags: []string{
			"#EXTM3U",
			"#EXT-X-VERSION:4",
			"#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=2231539,BANDWIDTH=2984657,CODECS=\"avc1.64001F,mp4a.40.2\",RESOLUTION=1280x720",
		},
		MediaPlaylists: []parse.MediaPlaylist{},
	}
	loader := NewDefaultLoader(original)
	actual, err := loader.LoadMasterPlaylist()
	if err != nil {
		t.Fatalf("error happened: %#v", err)
	}
	masterPlaylist := "#EXTM3U\n" +
		"#EXT-X-VERSION:4\n" +
		"#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=2231539,BANDWIDTH=2984657,CODECS=\"avc1.64001F,mp4a.40.2\",RESOLUTION=1280x720\n" +
		"0/playlist.m3u8\n"
	if masterPlaylist != string(actual) {
		t.Errorf("exspected: %v, actual: %v", masterPlaylist, string(actual))
	}
}
