package server

import "net/http"

type httpServer struct {
}

func New() *httpServer {
	return &httpServer{}
}

func (s *httpServer) Run(handler http.Handler) {
	runWithGracefulShutdown(handler)
}
