package main

import (
	"net/http"

	"github.com/OdaDaisuke/stream-go/internal/rtmp"
	"github.com/labstack/echo"
	"github.com/nareix/joy4/format"
)

func init() {
	format.RegisterAll()
}

func main() {
	rtmp_server := rtmp.NewRtmpServer()
	rtmp_server.Init()

	e := echo.New()
	e.GET("/streams", func(c echo.Context) error {
		// Get channels list
		return c.String(http.StatusOK, "")
	})
	e.PUT("/streams", func(c echo.Context) error {
		// Create New channel
		return c.String(http.StatusOK, "")
	})
	e.Logger.Fatal(e.Start(":8080"))

	rtmp_server.Start()
}
