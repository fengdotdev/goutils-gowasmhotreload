package server

import (
	"context"
	"net/http"
)

var _ ServerInterface = (*Server)(nil)

type Server struct {
	stop       chan bool
	lang       string
	httpServer *http.Server
}

func (s *Server) Start() error {

	return s.httpServer.ListenAndServe()
}

func (s *Server) StartNoBlock() {
	go s.Start()
}

func (s *Server) getIp() string {
	return "localhost:8080"
}

func (s *Server) Stop() error {
	return s.httpServer.Shutdown(context.Background())
}
