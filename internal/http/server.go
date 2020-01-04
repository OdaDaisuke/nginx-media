package http

import (
	"net/http"

	"github.com/OdaDaisuke/stream-go/internal/rtmp"
	"github.com/labstack/echo"
)

type HttpServer struct {
	e *echo.Echo
}

type PutStream struct {
	Name string `json:"name" form:"name" query:"name"`
}

func NewHttpServer(rtmpc *rtmp.RtmpCluster) *HttpServer {
	e := echo.New()
	e.GET("/streams", func(c echo.Context) error {
		return c.String(http.StatusOK, rtmpc.ListStreams())
	})
	e.PUT("/streams", func(c echo.Context) error {
		p := new(PutStream)
		if err := c.Bind(p); err != nil {
			return nil
		}
		rtmpc.AddStream(p.Name)
		return c.String(http.StatusOK, "")
	})
	return &HttpServer{e}
}

func (h *HttpServer) Run() {
	h.e.Logger.Fatal(h.e.Start(":8080"))
}
