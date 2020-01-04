package rtmp

import (
	"errors"
)

type IRtmpCluster interface {
	AddStream(string) *RtmpStream
	RemoveStream(string) error
	ListStreams() map[string]*RtmpStream
}

type RtmpCluster struct {
	streams map[string]*RtmpStream
}

func NewRtmpCluster() *RtmpCluster {
	return &RtmpCluster{}
}

func (r *RtmpCluster) Run() {
	r.AddStream("default")
}

func (r *RtmpCluster) ListStreams() map[string]*RtmpStream {
	return r.streams
}

func (r *RtmpCluster) AddStream(name string) {
	strm := NewRtmpStream()
	r.streams[name] = strm
}

func (r *RtmpCluster) RemoveStream(name string) error {
	_, ok := r.streams[name]
	if !ok {
		return errors.New("No stream found")
	}

	r.streams[name] = nil
	// TODO: Close connection
	return nil
}
