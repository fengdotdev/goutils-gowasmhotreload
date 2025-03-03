package server

import (
	"net/http"
)

func NewServer(addrs string, dir string) *Server {
	handler := http.FileServer(http.Dir(dir))
	return &Server{
		httpServer: &http.Server{
			Addr:    addrs, // ex. ":8080"
			Handler: handler,
		},
	}
}

func NewServerFromSettings(s Settings) *Server {
	return NewServer(s.GetPort(), s.GetDir())
}
