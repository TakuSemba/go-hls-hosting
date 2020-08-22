package load

type Loader interface {
	LoadMasterPlaylist() ([]byte, error)
	LoadMediaPlaylist(index int) ([]byte, error)
	LoadSegment(mediaPlaylistIndex, segmentIndex int) ([]byte, error)
}
