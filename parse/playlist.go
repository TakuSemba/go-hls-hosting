package parse

type ContainerFormat int

const (
	Ts ContainerFormat = iota
	Fmp4
)

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
	FileExtension         string
	ContainerFormat       ContainerFormat
}
