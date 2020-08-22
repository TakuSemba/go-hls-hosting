package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Parse Playlist
	parser := NewParser()
	original, err := parser.Parse("/Users/takusemba/Downloads/sample-bbb/output/master.m3u8")
	if err != nil {
		e.Logger.Fatal(err)
	}
	handler := NewHandler(original)

	// Routes
	e.GET("/hello", hello)
	e.GET("/vod/playlist.m3u8", handler.vodMasterPlaylist)
	e.GET("/vod/:index/playlist.m3u8", handler.vodMediaPlaylist)
	e.GET("/vod/:index/:segment", handler.vodSegments)
	e.GET("/live/playlist.m3u8", handler.liveMasterPlaylist)
	e.GET("/chase/playlist.m3u8", handler.chaseMasterPlaylist)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
