package main

type MasterPlaylist struct {
	Path           string
	Tags           []string
	MediaPlaylists []MediaPlaylist
}

type MediaPlaylist struct {
	Path     string
	Tags     []string
	Segments []Segment
}

type Segment struct {
	Path string
}

type Loader interface {
	loadMasterPlaylist() ([]byte, error)
	loadMediaPlaylist(index int) ([]byte, error)
	loadSegment(mediaPlaylistIndex, segmentIndex int) ([]byte, error)
}
