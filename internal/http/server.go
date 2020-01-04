package http

import (
	"net/http"

	"github.com/labstack/echo"
)

type HttpServer struct {
	e *echo.Echo
}

func NewHttpServer() *HttpServer {
	e := echo.New()
	e.GET("/streams", func(c echo.Context) error {
		// Get channels list
		return c.String(http.StatusOK, "")
	})
	e.PUT("/streams", func(c echo.Context) error {
		// Create New channel
		return c.String(http.StatusOK, "")
	})
	return &HttpServer{e}
}

func (h *HttpServer) Run() {
	h.e.Logger.Fatal(h.e.Start(":8080"))
}
