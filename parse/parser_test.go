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
	masterPlaylist, err := parser.ParseMasterPlaylist(testPath)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	tags := []string{
		"#EXTM3U",
		"#EXT-X-VERSION:4",
		"#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=2231539,BANDWIDTH=2984657,CODECS=\"avc1.64001F,mp4a.40.2\",RESOLUTION=1280x720",
	}
	if !reflect.DeepEqual(masterPlaylist.Tags, tags) {
		t.Errorf("exspected: %v, actual: %v", tags, masterPlaylist.Tags)
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
				"#EXT-X-TARGETDURATION:8\n" +
				"#EXT-X-MEDIA-SEQUENCE:0\n" +
				"#EXTINF:7.500000,\n" +
				"segment-0.ts\n" +
				"#EXTINF:6.916667,\n" +
				"segment-1.ts\n" +
				"#EXTINF:6.375000,\n" +
				"segment-2.ts\n" +
				"#EXTINF:7.291667,\n" +
				"segment-3.ts\n" +
				"#EXTINF:7.500000,\n" +
				"segment-4.ts\n" +
				"#EXTINF:7.500000,\n" +
				"segment-5.ts\n" +
				"#EXT-X-ENDLIST\n",
		),
	}
	parser := Parser{ReadFile: fileReader.FakeReadFile}
	mediaPlaylist, err := parser.ParseMediaPlaylist(testPath)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	tags := []string{
		"#EXTM3U",
		"#EXT-X-VERSION:4",
		"#EXT-X-PLAYLIST-TYPE:VOD",
		"#EXT-X-INDEPENDENT-SEGMENTS",
		"#EXT-X-TARGETDURATION:8",
		"#EXT-X-MEDIA-SEQUENCE:0",
		"#EXT-X-DISCONTINUITY-SEQUENCE:0",
		"#EXTINF:7.500000,",
		"#EXTINF:6.916667,",
		"#EXTINF:6.375000,",
		"#EXTINF:7.291667,",
		"#EXTINF:7.500000,",
		"#EXTINF:7.500000,",
		"#EXT-X-ENDLIST",
	}
	if !reflect.DeepEqual(mediaPlaylist.Tags, tags) {
		t.Errorf("exspected: %v, actual: %v", tags, mediaPlaylist.Tags)
	}
	segments := []Segment{
		{Path: "segment-0.ts", DurationMs: 7500.000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
		{Path: "segment-1.ts", DurationMs: 6916.667, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
		{Path: "segment-2.ts", DurationMs: 6375.000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
		{Path: "segment-3.ts", DurationMs: 7291.667, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
		{Path: "segment-4.ts", DurationMs: 7500.000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
		{Path: "segment-5.ts", DurationMs: 7500.000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
	}
	if !reflect.DeepEqual(mediaPlaylist.Segments, segments) {
		t.Errorf("exspected: %v, actual: %v", segments, mediaPlaylist.Segments)
	}
}
