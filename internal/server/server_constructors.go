package server

import (
	"net/http"

	"github.com/fengdotdev/goutils-gowasmhotreload/internal/settings"
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

func NewServerFromSettings(s *settings.SettingsTemplateInterface) *Server {
	return NewServer(":"+s(*sett).GetPort(), s.GetDir())
}
