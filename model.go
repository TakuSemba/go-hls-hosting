package main

type MasterPlaylist struct {
	Uri            string
	Tags           []string
	MediaPlaylists []MediaPlaylist
}

type MediaPlaylist struct {
	Uri      string
	Tags     []string
	Segments []Segment
}

type Segment struct {
	Uri string
}

type Loader interface {
	loadMasterPlaylist() ([]byte, error)
	loadMediaPlaylist(id int) ([]byte, error)
	loadSegment(id int) ([]byte, error)
}
