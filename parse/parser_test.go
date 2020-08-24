package parse

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

type FakeFileReader struct {
	expectedPath string
	ReadBytes    []byte
}

func (f *FakeFileReader) FakeReadFile(path string) ([]byte, error) {
	// return empty bytes if path is unexpected.
	if path != f.expectedPath {
		return []byte{}, nil
	}
	return f.ReadBytes, nil
}

func TestParseMasterPlaylist(t *testing.T) {
	readBytes, _ := ioutil.ReadFile("../testdata/master.m3u8")
	testPath := "testMasterPlaylist.m3u8"
	fileReader := FakeFileReader{
		expectedPath: testPath,
		ReadBytes:    readBytes,
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
	readBytes, _ := ioutil.ReadFile("../testdata/stream_ts.m3u8")
	testPath := "testMediaPlaylist.m3u8"
	fileReader := FakeFileReader{
		expectedPath: testPath,
		ReadBytes:    readBytes,
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
		"#EXT-X-ENDLIST",
	}
	totalDiscontinuityCount := 0
	totalDurationMs := float64(12000)
	segments := []Segment{
		{Path: "segment-0.ts", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
		{Path: "segment-1.ts", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
		{Path: "segment-2.ts", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
	}
	if assert.NoError(t, err) {
		assert.Equal(t, tags, actual.Tags)
		assert.Equal(t, totalDiscontinuityCount, actual.TotalDiscontinuityCount)
		assert.Equal(t, totalDurationMs, actual.TotalDurationMs)
		assert.Equal(t, segments, actual.Segments)
	}
}

func TestParseFmp4MediaPlaylist(t *testing.T) {
	readBytes, _ := ioutil.ReadFile("../testdata/stream_fmp4.m3u8")
	testPath := "testMediaPlaylist.m3u8"
	fileReader := FakeFileReader{
		expectedPath: testPath,
		ReadBytes:    readBytes,
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
		"#EXT-X-ENDLIST",
	}
	totalDiscontinuityCount := 0
	totalDurationMs := float64(12000)
	segments := []Segment{
		{Path: "segment-0.mp4", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: SegmentBySegment},
		{Path: "segment-1.mp4", DurationMs: 4000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: SegmentBySegment},
		{Path: "segment-2.mp4", DurationMs: 5000, DiscontinuitySequence: 0, FileExtension: ".mp4", ContainerFormat: Fmp4, RequestType: SegmentBySegment},
	}
	if assert.NoError(t, err) {
		assert.Equal(t, tags, actual.Tags)
		assert.Equal(t, totalDiscontinuityCount, actual.TotalDiscontinuityCount)
		assert.Equal(t, totalDurationMs, actual.TotalDurationMs)
		assert.Equal(t, segments, actual.Segments)
	}
}

func TestParseByteRangeMediaPlaylist(t *testing.T) {
	readBytes, _ := ioutil.ReadFile("../testdata/stream_byterange.m3u8")
	testPath := "testMediaPlaylist.m3u8"
	fileReader := FakeFileReader{
		expectedPath: testPath,
		ReadBytes:    readBytes,
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
		"#EXT-X-ENDLIST",
	}
	totalDiscontinuityCount := 0
	totalDurationMs := float64(12000)
	segments := []Segment{
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

func TestParseDiscontinuityMediaPlaylist(t *testing.T) {
	readBytes, _ := ioutil.ReadFile("../testdata/stream_discontinuity.m3u8")
	testPath := "testMediaPlaylist.m3u8"
	fileReader := FakeFileReader{
		expectedPath: testPath,
		ReadBytes:    readBytes,
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
		"#EXT-X-DISCONTINUITY",
		"#EXTINF:4,",
		"#EXTINF:5,",
		"#EXT-X-ENDLIST",
	}
	totalDiscontinuityCount := 1
	totalDurationMs := float64(12000)
	segments := []Segment{
		{Path: "segment-0.ts", DurationMs: 3000, DiscontinuitySequence: 0, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
		{Path: "segment-1.ts", DurationMs: 4000, DiscontinuitySequence: 1, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
		{Path: "segment-2.ts", DurationMs: 5000, DiscontinuitySequence: 1, FileExtension: ".ts", ContainerFormat: Ts, RequestType: SegmentBySegment},
	}
	if assert.NoError(t, err) {
		assert.Equal(t, tags, actual.Tags)
		assert.Equal(t, totalDiscontinuityCount, actual.TotalDiscontinuityCount)
		assert.Equal(t, totalDurationMs, actual.TotalDurationMs)
		assert.Equal(t, segments, actual.Segments)
	}
}
