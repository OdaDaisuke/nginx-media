package rtmp

import (
	"sync"

	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/av/pubsub"
	"github.com/nareix/joy4/format/rtmp"
)

type Channel struct {
	queue *pubsub.Queue
}

type RtmpServer struct {
	Server *rtmp.Server

	channels map[string]*Channel
	l        *sync.RWMutex
}

func NewRtmpServer() *RtmpServer {
	return &RtmpServer{
		Server: &rtmp.Server{},
		l:      &sync.RWMutex{},
	}
}

func (r *RtmpServer) Init() {
	r.Server.HandlePlay = func(conn *rtmp.Conn) {
		r.l.RLock()
		ch := r.channels[conn.URL.Path]
		r.l.RUnlock()
		if ch != nil {
			avutil.CopyFile(conn, ch.queue.Latest())
		}
	}

	r.Server.HandlePublish = func(conn *rtmp.Conn) {
		streams, _ := conn.Streams()
		r.l.Lock()
		ch := r.channels[conn.URL.Path]
		if ch == nil {
			ch = &Channel{}
			ch.queue = pubsub.NewQueue()
			ch.queue.WriteHeader(streams)
			r.channels[conn.URL.Path] = ch
		} else {
			ch = nil
		}
		r.l.Unlock()
		if ch == nil {
			return
		}

		avutil.CopyPackets(ch.queue, conn)
		r.l.Lock()
		delete(r.channels, conn.URL.Path)
		r.l.Unlock()
		ch.queue.Close()
	}
}

// Start streaming server
func (r *RtmpServer) Start() {
	r.Server.ListenAndServe()
}
