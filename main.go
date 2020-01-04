package main

import (
	"github.com/OdaDaisuke/stream-go/internal/http"
	"github.com/OdaDaisuke/stream-go/internal/rtmp"
	"github.com/nareix/joy4/format"
)

func init() {
	format.RegisterAll()
}

func main() {
	// API server for RTMP streams operation
	http_server := http.NewHttpServer()
	http_server.Run()

	// Expose RTMP Sterams
	rtmp_cluster := rtmp.NewRtmpCluster()
	rtmp_cluster.Run()
}
