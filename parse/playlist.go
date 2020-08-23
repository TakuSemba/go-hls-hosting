package parse

type MasterPlaylist struct {
	Path           string
	Tags           []string
	MediaPlaylists []MediaPlaylist
}

type MediaPlaylist struct {
	Path                    string
	Tags                    []string
	Segments                []Segment
	TotalDurationMs         float64
	TotalDiscontinuityCount int
}

type Segment struct {
	Path                  string
	DurationMs            float64
	DiscontinuitySequence int
}
