package handle

import (
	"github.com/TakuSemba/go-media-hosting/load"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
	masterPlaylist, err := h.VodLoader.LoadMasterPlaylist()
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MasterPlaylist.")
	}
	return c.Blob(http.StatusOK, "application/x-mpegURL", masterPlaylist)
}

func (h *Handler) VodMediaPlaylist(c echo.Context) error {
	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	mediaPlaylist, err := h.VodLoader.LoadMediaPlaylist(index)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	return c.Blob(http.StatusOK, "application/x-mpegURL", mediaPlaylist)
}

func (h *Handler) VodSegments(c echo.Context) error {
	mediaPlaylistIndex, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	segmentName := c.Param("segment")
	segmentIndex, err := strconv.Atoi(segmentName[0 : len(segmentName)-3])
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	segment, err := h.VodLoader.LoadSegment(mediaPlaylistIndex, segmentIndex)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load Segment.")
	}
	// TODO Change Content-Type
	return c.Blob(http.StatusOK, "video/MP2T", segment)
}

func (h *Handler) LiveMasterPlaylist(c echo.Context) error {
	masterPlaylist, err := h.LiveLoader.LoadMasterPlaylist()
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MasterPlaylist.")
	}
	return c.Blob(http.StatusOK, "application/x-mpegURL", masterPlaylist)
}

func (h *Handler) LiveMediaPlaylist(c echo.Context) error {
	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	mediaPlaylist, err := h.LiveLoader.LoadMediaPlaylist(index)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	return c.Blob(http.StatusOK, "application/x-mpegURL", mediaPlaylist)
}

func (h *Handler) LiveSegments(c echo.Context) error {
	mediaPlaylistIndex, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	segmentName := c.Param("segment")
	segmentIndex, err := strconv.Atoi(segmentName[0 : len(segmentName)-3])
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	segment, err := h.LiveLoader.LoadSegment(mediaPlaylistIndex, segmentIndex)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load Segment.")
	}
	// TODO Change Content-Type
	return c.Blob(http.StatusOK, "video/MP2T", segment)
}

func (h *Handler) ChaseMasterPlaylist(c echo.Context) error {
	masterPlaylist, err := h.ChaseLoader.LoadMasterPlaylist()
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MasterPlaylist.")
	}
	return c.Blob(http.StatusOK, "application/x-mpegURL", masterPlaylist)
}

func (h *Handler) ChaseMediaPlaylist(c echo.Context) error {
	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	mediaPlaylist, err := h.ChaseLoader.LoadMediaPlaylist(index)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	return c.Blob(http.StatusOK, "application/x-mpegURL", mediaPlaylist)
}

func (h *Handler) ChaseSegments(c echo.Context) error {
	mediaPlaylistIndex, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	segmentName := c.Param("segment")
	segmentIndex, err := strconv.Atoi(segmentName[0 : len(segmentName)-3])
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	segment, err := h.ChaseLoader.LoadSegment(mediaPlaylistIndex, segmentIndex)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load Segment.")
	}
	// TODO Change Content-Type
	return c.Blob(http.StatusOK, "video/MP2T", segment)
}
