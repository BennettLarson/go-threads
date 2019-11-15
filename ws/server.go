package ws

import (
	"context"
	"net/http"
	"time"

	logging "github.com/ipfs/go-log"
	tserv "github.com/textileio/go-textile-core/threadservice"
)

var log = logging.Logger("ws")

// Server wraps a connection hub and http server.
type Server struct {
	service tserv.Threadservice
	hub     *Hub
	s       *http.Server
}

// NewServer returns a web socket server.
func NewServer(ctx context.Context, service tserv.Threadservice, addr string) *Server {
	s := &Server{
		hub: newHub(ctx, service),
	}
	go s.hub.run()

	s.s = &http.Server{
		Addr: addr,
	}
	s.s.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serveWs(s.hub, w, r)
	})

	errc := make(chan error)
	go func() {
		errc <- s.s.ListenAndServe()
		close(errc)
	}()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err, ok := <-errc:
				if err != nil && err != http.ErrServerClosed {
					log.Errorf("ws server error: %s", err)
				}
				if !ok {
					log.Info("ws server was shutdown")
					return
				}
			}
		}
	}()

	sub := service.Subscribe()
	go func() {
		for {
			select {
			case rec, ok := <-sub.Channel():
				if !ok {
					return
				}
				s.hub.broadcast <- rec
			}
		}
	}()

	log.Infof("ws server listening at %s", s.s.Addr)

	return s
}

// Close the server.
func (s *Server) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := s.s.Shutdown(ctx); err != nil {
		log.Errorf("error shutting down ws server: %s", err)
	}
}
