package load

import (
	"github.com/TakuSemba/go-media-hosting/parse"
	"testing"
	"time"
)

func TestChaseLoadMediaPlaylist(t *testing.T) {
	original := parse.MasterPlaylist{
		MediaPlaylists: []parse.MediaPlaylist{
			{
				Path: "testMediaPlaylist.m3u8",
				Tags: []string{
					"#EXTM3U",
					"#EXT-X-VERSION:4",
					"#EXT-X-PLAYLIST-TYPE:VOD",
					"#EXT-X-INDEPENDENT-SEGMENTS",
					"#EXT-X-TARGETDURATION:5",
					"#EXT-X-MEDIA-SEQUENCE:0",
					"#EXT-X-DISCONTINUITY-SEQUENCE:0",
					"#EXTINF:3,",
					"#EXTINF:4,",
					"#EXTINF:5,",
					"#EXTINF:3,",
					"#EXTINF:4,",
					"#EXTINF:5,",
					"#EXTINF:3,",
					"#EXTINF:4,",
					"#EXTINF:5,",
					"#EXT-X-ENDLIST",
				},
				Segments: []parse.Segment{
					{Path: "segment-0.ts", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
					{Path: "segment-1.ts", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
					{Path: "segment-2.ts", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
					{Path: "segment-3.ts", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
					{Path: "segment-4.ts", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
					{Path: "segment-5.ts", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
					{Path: "segment-6.ts", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
					{Path: "segment-7.ts", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
					{Path: "segment-8.ts", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: parse.Ts, RequestType: parse.SegmentBySegment},
				},
				TotalDurationMs:         36 * 1000,
				TotalDiscontinuityCount: 0,
			},
		},
	}
	loader := NewChaseLoader(original)
	// 16 seconds before
	loader.StartedAt = time.Now().Add(time.Duration(-16 * 1000 * 1000 * 1000))
	actual, err := loader.LoadMediaPlaylist(0)
	if err != nil {
		t.Fatalf("error happened: %#v", err)
	}
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
	if mediaPlaylist != string(actual) {
		t.Errorf("exspected: %v, actual: %v", mediaPlaylist, string(actual))
	}
}
