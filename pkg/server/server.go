package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutDownTimeout time.Duration
}

const (
	_defaultShutdownTimeout = 5 * time.Second
)

func New(handler http.Handler, address string) *Server {
	httpServer := &http.Server{
		Handler: handler,
		Addr:    address,
	}
	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutDownTimeout: _defaultShutdownTimeout,
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutDownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}
