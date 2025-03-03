package server

import (
	"fmt"
	"net/http"

	"github.com/fengdotdev/goutils-gowasmhotreload/internal/msg"
)

var _ ServerInterface = (*Server)(nil)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Start() error {

	host := fmt.Sprintf("http://%s", getIp())
	fmt.Println(msg.MSG.ServerRunningAt, host+s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func getIp() string {
	return "localhost"
}
