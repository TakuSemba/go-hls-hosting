package parse

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	FakeMasterPlaylist = "#EXTM3U\n" +
		"#EXT-X-VERSION:4\n" +
		"#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=2231539,BANDWIDTH=2984657,CODECS=\"avc1.64001F,mp4a.40.2\",RESOLUTION=1280x720\n" +
		"stream-1/stream.m3u8\n"

	FakeTsMediaPlaylist = "#EXTM3U\n" +
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
		"#EXT-X-ENDLIST\n"

	FakeFmp4MediaPlaylist = "#EXTM3U\n" +
		"#EXT-X-VERSION:4\n" +
		"#EXT-X-PLAYLIST-TYPE:VOD\n" +
		"#EXT-X-INDEPENDENT-SEGMENTS\n" +
		"#EXT-X-TARGETDURATION:5\n" +
		"#EXT-X-MEDIA-SEQUENCE:0\n" +
		"#EXTINF:3,\n" +
		"segment-0.mp4\n" +
		"#EXTINF:4,\n" +
		"segment-1.mp4\n" +
		"#EXTINF:5,\n" +
		"segment-2.mp4\n" +
		"#EXTINF:3,\n" +
		"segment-3.mp4\n" +
		"#EXTINF:4,\n" +
		"segment-4.mp4\n" +
		"#EXTINF:5,\n" +
		"segment-5.mp4\n" +
		"#EXTINF:3,\n" +
		"segment-6.mp4\n" +
		"#EXTINF:4,\n" +
		"segment-7.mp4\n" +
		"#EXTINF:5,\n" +
		"segment-8.mp4\n" +
		"#EXT-X-ENDLIST\n"

	FakeByteRangeMediaPlaylist = "#EXTM3U\n" +
		"#EXT-X-VERSION:4\n" +
		"#EXT-X-PLAYLIST-TYPE:VOD\n" +
		"#EXT-X-INDEPENDENT-SEGMENTS\n" +
		"#EXT-X-TARGETDURATION:5\n" +
		"#EXT-X-MEDIA-SEQUENCE:0\n" +
		"#EXTINF:3,\n" +
		"#EXT-X-BYTERANGE:300000@100\n" +
		"segment-0.mp4\n" +
		"#EXTINF:4,\n" +
		"#EXT-X-BYTERANGE:400000@300100\n" +
		"segment-0.mp4\n" +
		"#EXTINF:5,\n" +
		"#EXT-X-BYTERANGE:500000@700100\n" +
		"segment-0.mp4\n" +
		"#EXTINF:3,\n" +
		"#EXT-X-BYTERANGE:300000@1200100\n" +
		"segment-0.mp4\n" +
		"#EXTINF:4,\n" +
		"#EXT-X-BYTERANGE:400000@1500100\n" +
		"segment-0.mp4\n" +
		"#EXTINF:5,\n" +
		"#EXT-X-BYTERANGE:500000@1900100\n" +
		"segment-0.mp4\n" +
		"#EXTINF:3,\n" +
		"#EXT-X-BYTERANGE:300000@2400100\n" +
		"segment-0.mp4\n" +
		"#EXTINF:4,\n" +
		"#EXT-X-BYTERANGE:400000@2700100\n" +
		"segment-0.mp4\n" +
		"#EXTINF:5,\n" +
		"#EXT-X-BYTERANGE:500000@3100100\n" +
		"segment-0.mp4\n" +
		"#EXT-X-ENDLIST\n"
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
		ReadBytes:    []byte(FakeMasterPlaylist),
	}
	parser := Parser{ReadFile: fileReader.FakeReadFile}
	actual, err := parser.ParseMasterPlaylist(testPath)
	tags := []string{
		"#EXTM3U",
		"#EXT-X-VERSION:4",
		"#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=2231539,BANDWIDTH=2984657,CODECS=\"avc1.64001F,mp4a.40.2\",RESOLUTION=1280x720",
	}
	if assert.NoError(t, err) {
		assert.Equal(t, tags, actual.Tags)
	}
}

func TestParseTsMediaPlaylist(t *testing.T) {
	testPath := "testMediaPlaylist.m3u8"
	fileReader := FakeFileReader{
		expectedPath: testPath,
		ReadBytes:    []byte(FakeTsMediaPlaylist),
	}
	parser := Parser{ReadFile: fileReader.FakeReadFile}
	actual, err := parser.ParseMediaPlaylist(testPath)
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
	totalDiscontinuityCount := 0
	totalDurationMs := float64(36000)
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
	if assert.NoError(t, err) {
		assert.Equal(t, tags, actual.Tags)
		assert.Equal(t, totalDiscontinuityCount, actual.TotalDiscontinuityCount)
		assert.Equal(t, totalDurationMs, actual.TotalDurationMs)
		assert.Equal(t, segments, actual.Segments)
	}
}

func TestParseFmp4MediaPlaylist(t *testing.T) {
	testPath := "testMediaPlaylist.m3u8"
	fileReader := FakeFileReader{
		expectedPath: testPath,
		ReadBytes:    []byte(FakeFmp4MediaPlaylist),
	}
	parser := Parser{ReadFile: fileReader.FakeReadFile}
	actual, err := parser.ParseMediaPlaylist(testPath)
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
	totalDiscontinuityCount := 0
	totalDurationMs := float64(36000)
	segments := []Segment{
		{Path: "segment-0.mp4", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: SegmentBySegment},
		{Path: "segment-1.mp4", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: SegmentBySegment},
		{Path: "segment-2.mp4", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: SegmentBySegment},
		{Path: "segment-3.mp4", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: SegmentBySegment},
		{Path: "segment-4.mp4", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: SegmentBySegment},
		{Path: "segment-5.mp4", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: SegmentBySegment},
		{Path: "segment-6.mp4", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: SegmentBySegment},
		{Path: "segment-7.mp4", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: SegmentBySegment},
		{Path: "segment-8.mp4", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: SegmentBySegment},
	}
	if assert.NoError(t, err) {
		assert.Equal(t, tags, actual.Tags)
		assert.Equal(t, totalDiscontinuityCount, actual.TotalDiscontinuityCount)
		assert.Equal(t, totalDurationMs, actual.TotalDurationMs)
		assert.Equal(t, segments, actual.Segments)
	}
}

func TestParseByteRangeMediaPlaylist(t *testing.T) {
	testPath := "testMediaPlaylist.m3u8"
	fileReader := FakeFileReader{
		expectedPath: testPath,
		ReadBytes:    []byte(FakeByteRangeMediaPlaylist),
	}
	parser := Parser{ReadFile: fileReader.FakeReadFile}
	actual, err := parser.ParseMediaPlaylist(testPath)
	tags := []string{
		"#EXTM3U",
		"#EXT-X-VERSION:4",
		"#EXT-X-PLAYLIST-TYPE:VOD",
		"#EXT-X-INDEPENDENT-SEGMENTS",
		"#EXT-X-TARGETDURATION:5",
		"#EXT-X-MEDIA-SEQUENCE:0",
		"#EXT-X-DISCONTINUITY-SEQUENCE:0",
		"#EXTINF:3,",
		"#EXT-X-BYTERANGE:300000@100",
		"#EXTINF:4,",
		"#EXT-X-BYTERANGE:400000@300100",
		"#EXTINF:5,",
		"#EXT-X-BYTERANGE:500000@700100",
		"#EXTINF:3,",
		"#EXT-X-BYTERANGE:300000@1200100",
		"#EXTINF:4,",
		"#EXT-X-BYTERANGE:400000@1500100",
		"#EXTINF:5,",
		"#EXT-X-BYTERANGE:500000@1900100",
		"#EXTINF:3,",
		"#EXT-X-BYTERANGE:300000@2400100",
		"#EXTINF:4,",
		"#EXT-X-BYTERANGE:400000@2700100",
		"#EXTINF:5,",
		"#EXT-X-BYTERANGE:500000@3100100",
		"#EXT-X-ENDLIST",
	}
	totalDiscontinuityCount := 0
	totalDurationMs := float64(36000)
	segments := []Segment{
		{Path: "segment-0.mp4", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: ByteRange},
		{Path: "segment-0.mp4", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: ByteRange},
		{Path: "segment-0.mp4", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: ByteRange},
		{Path: "segment-0.mp4", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: ByteRange},
		{Path: "segment-0.mp4", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: ByteRange},
		{Path: "segment-0.mp4", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: ByteRange},
		{Path: "segment-0.mp4", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: ByteRange},
		{Path: "segment-0.mp4", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: ByteRange},
		{Path: "segment-0.mp4", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: ByteRange},
	}
	if assert.NoError(t, err) {
		assert.Equal(t, tags, actual.Tags)
		assert.Equal(t, totalDiscontinuityCount, actual.TotalDiscontinuityCount)
		assert.Equal(t, totalDurationMs, actual.TotalDurationMs)
		assert.Equal(t, segments, actual.Segments)
	}
}
