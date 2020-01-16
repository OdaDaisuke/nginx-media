package rtmp

import (
	"errors"
	"fmt"
)

type IRtmpCluster interface {
	AddStream(string) *RtmpStream
	RemoveStream(string) error
	ListStreams() string
}

type RtmpCluster struct {
	streams map[string]*RtmpStream
}

func NewRtmpCluster() *RtmpCluster {
	return &RtmpCluster{
		streams: make(map[string]*RtmpStream),
	}
}

func (r *RtmpCluster) Run() {
	r.AddStream("default")
}

func (r *RtmpCluster) ListStreams() string {
	return fmt.Sprintf("%s", r.streams)
}

func (r *RtmpCluster) AddStream(name string) {
	strm := NewRtmpStream()
	strm.Run()
	r.streams[name] = strm
}

func (r *RtmpCluster) RemoveStream(name string) error {
	_, ok := r.streams[name]
	if !ok {
		return errors.New("No stream found")
	}

	delete(r.streams, name)
	// TODO: Close connection
	return nil
}
