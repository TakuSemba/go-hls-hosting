package handle

import (
	"github.com/TakuSemba/go-media-hosting/load"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

const (
	TsFileExtension        = ".ts"
	Mp4FileExtension       = ".mp4"
	M4FileExtensionPrefix  = ".m4"
	Mp4FileExtensionPrefix = ".mp4"
	CmfFileExtensionPrefix = ".cmf"
)

const (
	ContentTypeMpegUrl = "application/x-mpegURL"
	ContentTypeMpeg2Ts = "video/MP2T"
	ContentTypeMP4     = "video/mp4"
)

type Handler struct {
	VodLoader   load.Loader
	LiveLoader  load.Loader
	ChaseLoader load.Loader
}

func NewHandler(vodLoader load.VodLoader, liveLoader load.LiveLoader, chaseLoader load.ChaseLoader) Handler {
	return Handler{
		VodLoader:   &vodLoader,
		LiveLoader:  &liveLoader,
		ChaseLoader: &chaseLoader,
	}
}

func (h *Handler) VodMasterPlaylist(c echo.Context) error {
	return h.loadMasterPlaylist(h.VodLoader, c)
}

func (h *Handler) VodMediaPlaylist(c echo.Context) error {
	return h.loadMediaPlaylist(h.VodLoader, c)
}

func (h *Handler) VodSegments(c echo.Context) error {
	return h.loadSegment(h.ChaseLoader, c)
}

func (h *Handler) LiveMasterPlaylist(c echo.Context) error {
	return h.loadMasterPlaylist(h.LiveLoader, c)
}

func (h *Handler) LiveMediaPlaylist(c echo.Context) error {
	return h.loadMediaPlaylist(h.LiveLoader, c)
}

func (h *Handler) LiveSegments(c echo.Context) error {
	return h.loadSegment(h.ChaseLoader, c)
}

func (h *Handler) ChaseMasterPlaylist(c echo.Context) error {
	return h.loadMasterPlaylist(h.ChaseLoader, c)
}

func (h *Handler) ChaseMediaPlaylist(c echo.Context) error {
	return h.loadMediaPlaylist(h.ChaseLoader, c)
}

func (h *Handler) ChaseSegments(c echo.Context) error {
	return h.loadSegment(h.ChaseLoader, c)
}

func (h *Handler) loadMasterPlaylist(loader load.Loader, c echo.Context) error {
	masterPlaylist, err := loader.LoadMasterPlaylist()
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MasterPlaylist.")
	}
	return c.Blob(http.StatusOK, ContentTypeMpegUrl, masterPlaylist)
}

func (h *Handler) loadMediaPlaylist(loader load.Loader, c echo.Context) error {
	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	mediaPlaylist, err := loader.LoadMediaPlaylist(index)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	return c.Blob(http.StatusOK, ContentTypeMpegUrl, mediaPlaylist)
}

func (h *Handler) loadSegment(loader load.Loader, c echo.Context) error {
	mediaPlaylistIndex, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	segmentName := c.Param("segment")
	segmentIndex, err := strconv.Atoi(segmentName[0 : len(segmentName)-3])
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	segment, err := loader.LoadSegment(mediaPlaylistIndex, segmentIndex)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load Segment.")
	}

	var contentType string
	switch {
	case strings.HasSuffix(segmentName, TsFileExtension):
		contentType = ContentTypeMpeg2Ts
	case strings.HasSuffix(segmentName, Mp4FileExtension):
		contentType = ContentTypeMP4
	case strings.HasPrefix(segmentName[len(segmentName)-4:], M4FileExtensionPrefix):
		contentType = ContentTypeMP4
	case strings.HasPrefix(segmentName[len(segmentName)-5:], Mp4FileExtensionPrefix):
		contentType = ContentTypeMP4
	case strings.HasPrefix(segmentName[len(segmentName)-5:], CmfFileExtensionPrefix):
		contentType = ContentTypeMP4
	default:
		return c.String(http.StatusBadRequest, "failed to load Segment.")
	}

	return c.Blob(http.StatusOK, contentType, segment)
}
