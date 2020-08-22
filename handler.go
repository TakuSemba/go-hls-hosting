package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	VodLoader   Loader
	LiveLoader  Loader
	ChaseLoader Loader
}

func (h *Handler) vodMasterPlaylist(c echo.Context) error {
	masterPlaylist, err := h.VodLoader.loadMasterPlaylist()
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MasterPlaylist.")
	}
	return c.Blob(http.StatusOK, "application/x-mpegURL", masterPlaylist)
}

func (h *Handler) vodMediaPlaylist(c echo.Context) error {
	mediaPlaylist, err := h.VodLoader.loadMediaPlaylist(0)
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MediaPlaylist.")
	}
	return c.Blob(http.StatusOK, "application/x-mpegURL", mediaPlaylist)
}

func (h *Handler) vodSegments(c echo.Context) error {
	segment, err := h.VodLoader.loadSegment(0)
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

func (h *Handler) chaseMasterPlaylist(c echo.Context) error {
	masterPlaylist, err := h.ChaseLoader.loadMasterPlaylist()
	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to load MasterPlaylist.")
	}
	return c.Blob(http.StatusOK, "application/x-mpegURL", masterPlaylist)
}
