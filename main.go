package main

import (
	"github.com/TakuSemba/go-media-hosting/handle"
	"github.com/TakuSemba/go-media-hosting/load"
	"github.com/TakuSemba/go-media-hosting/parse"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Parse Playlist
	parser := parse.NewParser()
	original, err := parser.Parse("/Users/takusemba/Downloads/sample-bbb/output/master.m3u8")
	if err != nil {
		e.Logger.Fatal(err)
	}
	vodLoader := load.NewVodLoader(original)
	liveLoader := load.NewLiveLoader(original)
	chaseLoader := load.NewChaseLoader(original)
	handler := handle.NewHandler(vodLoader, liveLoader, chaseLoader)

	// Routes
	e.GET("/vod/playlist.m3u8", handler.VodMasterPlaylist)
	e.GET("/vod/:index/playlist.m3u8", handler.VodMediaPlaylist)
	e.GET("/vod/:index/:segment", handler.VodSegment)
	e.GET("/live/playlist.m3u8", handler.LiveMasterPlaylist)
	e.GET("/live/:index/playlist.m3u8", handler.LiveMediaPlaylist)
	e.GET("/live/:index/:segment", handler.LiveSegment)
	e.GET("/chase/playlist.m3u8", handler.ChaseMasterPlaylist)
	e.GET("/chase/:index/playlist.m3u8", handler.ChaseMediaPlaylist)
	e.GET("/chase/:index/:segment", handler.ChaseSegment)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
