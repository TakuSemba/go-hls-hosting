package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	VodLoader   Loader
	LiveLoader  Loader
	ChaseLoader Loader
}

func NewHandler(original MasterPlaylist) Handler {
	return Handler{
		VodLoader:   &VodLoader{MasterPlaylist: original},
		LiveLoader:  &LiveLoader{MasterPlaylist: original},
		ChaseLoader: &ChaseLoader{MasterPlaylist: original},
	}
}

func (h *Handler) vodMasterPlaylist(c echo.Context) error {
	masterPlaylist, err := h.VodLoader.loadMasterPlaylist()
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MasterPlaylist.")
	}
	return c.Blob(http.StatusOK, "application/x-mpegURL", masterPlaylist)
}

func (h *Handler) vodMediaPlaylist(c echo.Context) error {
	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	mediaPlaylist, err := h.VodLoader.loadMediaPlaylist(index)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	return c.Blob(http.StatusOK, "application/x-mpegURL", mediaPlaylist)
}

func (h *Handler) vodSegments(c echo.Context) error {
	mediaPlaylistIndex, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	segmentName := c.Param("segment")
	segmentIndex, err := strconv.Atoi(segmentName[0 : len(segmentName)-3])
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	segment, err := h.VodLoader.loadSegment(mediaPlaylistIndex, segmentIndex)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load Segment.")
	}
	// TODO Change Content-Type
	return c.Blob(http.StatusOK, "video/MP2T", segment)
}

func (h *Handler) liveMasterPlaylist(c echo.Context) error {
	masterPlaylist, err := h.LiveLoader.loadMasterPlaylist()
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MasterPlaylist.")
	}
	return c.Blob(http.StatusOK, "application/x-mpegURL", masterPlaylist)
}

func (h *Handler) liveMediaPlaylist(c echo.Context) error {
	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	mediaPlaylist, err := h.LiveLoader.loadMediaPlaylist(index)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	return c.Blob(http.StatusOK, "application/x-mpegURL", mediaPlaylist)
}

func (h *Handler) liveSegments(c echo.Context) error {
	mediaPlaylistIndex, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	segmentName := c.Param("segment")
	segmentIndex, err := strconv.Atoi(segmentName[0 : len(segmentName)-3])
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	segment, err := h.LiveLoader.loadSegment(mediaPlaylistIndex, segmentIndex)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load Segment.")
	}
	// TODO Change Content-Type
	return c.Blob(http.StatusOK, "video/MP2T", segment)
}

func (h *Handler) chaseMasterPlaylist(c echo.Context) error {
	masterPlaylist, err := h.ChaseLoader.loadMasterPlaylist()
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MasterPlaylist.")
	}
	return c.Blob(http.StatusOK, "application/x-mpegURL", masterPlaylist)
}

func (h *Handler) chaseMediaPlaylist(c echo.Context) error {
	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	mediaPlaylist, err := h.ChaseLoader.loadMediaPlaylist(index)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	return c.Blob(http.StatusOK, "application/x-mpegURL", mediaPlaylist)
}

func (h *Handler) chaseSegments(c echo.Context) error {
	mediaPlaylistIndex, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	segmentName := c.Param("segment")
	segmentIndex, err := strconv.Atoi(segmentName[0 : len(segmentName)-3])
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	segment, err := h.ChaseLoader.loadSegment(mediaPlaylistIndex, segmentIndex)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load Segment.")
	}
	// TODO Change Content-Type
	return c.Blob(http.StatusOK, "video/MP2T", segment)
}
