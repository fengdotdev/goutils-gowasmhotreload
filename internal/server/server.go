package server

import (
	"fmt"
	"net/http"

	"github.com/fengdotdev/goutils-gowasmhotreload/internal/msg"
	"github.com/fengdotdev/goutils-gowasmhotreload/internal/settings"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	handler := http.FileServer(http.Dir(settings.Settings.Dir))
	return &Server{
		httpServer: &http.Server{
			Addr: settings.Settings.Port,
			Handler: handler,
		},
	}
}


func (s *Server) Start() error {
	fmt.Println(msg.MSG.ServerRunningAt, "http://localhost"+ settings.Settings.Port)
	return s.httpServer.ListenAndServe()
}