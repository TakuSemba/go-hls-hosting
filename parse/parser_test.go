package parse

import (
	"reflect"
	"testing"
)

type FakeFileReader struct {
	expectedPath string
	ReadBytes    []byte
}

func (f *FakeFileReader) FakeReadFile(path string) ([]byte, error) {
	if path != f.expectedPath {
		return []byte{}, nil
	}
	return f.ReadBytes, nil
}

func TestParseMasterPlaylist(t *testing.T) {
	testPath := "testMasterPlaylist.m3u8"
	fileReader := FakeFileReader{
		expectedPath: testPath,
		ReadBytes: []byte(
			"#EXTM3U\n" +
				"#EXT-X-VERSION:4\n" +
				"#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=2231539,BANDWIDTH=2984657,CODECS=\"avc1.64001F,mp4a.40.2\",RESOLUTION=1280x720\n" +
				"media-1/stream.m3u8\n",
		),
	}
	parser := Parser{ReadFile: fileReader.FakeReadFile}
	actual, err := parser.ParseMasterPlaylist(testPath)
	if err != nil {
		t.Fatalf("error happened: %#v", err)
	}
	tags := []string{
		"#EXTM3U",
		"#EXT-X-VERSION:4",
		"#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=2231539,BANDWIDTH=2984657,CODECS=\"avc1.64001F,mp4a.40.2\",RESOLUTION=1280x720",
	}
	if !reflect.DeepEqual(tags, actual.Tags) {
		t.Errorf("exspected: %v, actual: %v", tags, actual.Tags)
	}
}

func TestParseMediaPlaylist(t *testing.T) {
	testPath := "testMediaPlaylist.m3u8"
	fileReader := FakeFileReader{
		expectedPath: testPath,
		ReadBytes: []byte(
			"#EXTM3U\n" +
				"#EXT-X-VERSION:4\n" +
				"#EXT-X-PLAYLIST-TYPE:VOD\n" +
				"#EXT-X-INDEPENDENT-SEGMENTS\n" +
				"#EXT-X-TARGETDURATION:5\n" +
				"#EXT-X-MEDIA-SEQUENCE:0\n" +
				"#EXTINF:3,\n" +
				"segment-0.ts\n" +
				"#EXTINF:4,\n" +
				"segment-1.ts\n" +
				"#EXTINF:5,\n" +
				"segment-2.ts\n" +
				"#EXTINF:3,\n" +
				"segment-3.ts\n" +
				"#EXTINF:4,\n" +
				"segment-4.ts\n" +
				"#EXTINF:5,\n" +
				"segment-5.ts\n" +
				"#EXTINF:3,\n" +
				"segment-6.ts\n" +
				"#EXTINF:4,\n" +
				"segment-7.ts\n" +
				"#EXTINF:5,\n" +
				"segment-8.ts\n" +
				"#EXT-X-ENDLIST\n",
		),
	}
	parser := Parser{ReadFile: fileReader.FakeReadFile}
	actual, err := parser.ParseMediaPlaylist(testPath)
	if err != nil {
		t.Fatalf("error happened: %#v", err)
	}
	tags := []string{
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
	}
	if !reflect.DeepEqual(tags, actual.Tags) {
		t.Errorf("exspected: %v, actual: %v", tags, actual.Tags)
	}
	totalDiscontinuityCount := 0
	if totalDiscontinuityCount != actual.TotalDiscontinuityCount {
		t.Errorf("exspected: %v, actual: %v", totalDiscontinuityCount, actual.TotalDiscontinuityCount)
	}
	totalDurationMs := float64(36000)
	if totalDurationMs != actual.TotalDurationMs {
		t.Errorf("exspected: %v, actual: %v", totalDurationMs, actual.TotalDurationMs)
	}
	segments := []Segment{
		{Path: "segment-0.ts", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
		{Path: "segment-1.ts", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
		{Path: "segment-2.ts", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
		{Path: "segment-3.ts", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
		{Path: "segment-4.ts", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
		{Path: "segment-5.ts", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
		{Path: "segment-6.ts", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
		{Path: "segment-7.ts", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
		{Path: "segment-8.ts", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
	}
	if !reflect.DeepEqual(segments, actual.Segments) {
		t.Errorf("exspected: %v, actual: %v", segments, actual.Segments)
	}
}
