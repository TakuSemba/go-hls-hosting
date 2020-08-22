package parse

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