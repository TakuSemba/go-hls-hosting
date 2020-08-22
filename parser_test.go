package main

import (
	"reflect"
	"testing"
)

func TestParseMasterPlaylist(t *testing.T) {
	data := "#EXTM3U\n" +
		"#EXT-X-VERSION:4\n" +
		"#EXT-X-STREAM-INF:AVERAGE-BANDWIDTH=2231539,BANDWIDTH=2984657,CODECS=\"avc1.64001F,mp4a.40.2\",RESOLUTION=1280x720\n" +
		"media-1/stream.m3u8\n"
	parser := Parser{}
	masterPlaylist, err := parser.ParseMasterPlaylist([]byte(data))
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
	data := "#EXTM3U\n" +
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
		"#EXT-X-ENDLIST\n"

	parser := Parser{}
	mediaPlaylist, err := parser.ParseMediaPlaylist([]byte(data))
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
		{Uri: "segment-0.ts"},
		{Uri: "segment-1.ts"},
		{Uri: "segment-2.ts"},
		{Uri: "segment-3.ts"},
		{Uri: "segment-4.ts"},
		{Uri: "segment-5.ts"},
	}
	if !reflect.DeepEqual(mediaPlaylist.Segments, segments) {
		t.Errorf("exspected: %v, actual: %v", segments, mediaPlaylist.Segments)
	}
}
