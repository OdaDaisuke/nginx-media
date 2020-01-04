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
	// Expose RTMP Sterams
	rtmp_cluster := rtmp.NewRtmpCluster()
	rtmp_cluster.Run()

	// API server for RTMP streams operation
	http_server := http.NewHttpServer(rtmp_cluster)
	http_server.Run()
}
