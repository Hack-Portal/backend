package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Hack-Portal/backend/cmd/config"
)

const (
	ShutdownTimeout = 10 * time.Second
)

func (s *httpServer) run(handler http.Handler) {
	s.srv = &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Config.Server.Addr),
		Handler: handler,
	}

	go func() {
		log.Printf("server lisening on %s\n", config.Config.Server.Addr)
		if err := s.srv.ListenAndServe(); err != nil || err != http.ErrServerClosed {
			log.Fatalf("Listen and serve failed: %+v\n", err)
		}
	}()
}
