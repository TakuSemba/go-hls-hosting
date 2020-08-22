package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {

	parser := Parser{}
	original, err := parser.Parse("/Users/takusemba/Downloads/sample-bbb/output/master.m3u8")
	if err != nil {
		return
	}

	handler := Handler{
		VodLoader:   &VodLoader{MasterPlaylist: original},
		LiveLoader:  &LiveLoader{MasterPlaylist: original},
		ChaseLoader: &ChaseLoader{MasterPlaylist: original},
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", hello)
	e.GET("/vod/playlist.m3u8", handler.vodMasterPlaylist)
	e.GET("/vod/:path/playlist.m3u8", handler.vodMediaPlaylist)
	e.GET("/vod/:path/*", handler.vodSegments)
	e.GET("/live/playlist.m3u8", handler.liveMasterPlaylist)
	e.GET("/chase/playlist.m3u8", handler.chaseMasterPlaylist)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
