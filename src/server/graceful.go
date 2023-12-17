package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hackhack-Geek-vol6/backend/cmd/config"
)

const (
	ShutdownTimeout = 10 * time.Second
)

func runWithGracefulShutdown(handler http.Handler) {
	srv := &http.Server{
		Addr:    config.Config.Server.Addr,
		Handler: handler,
	}

	go func() {
		log.Printf("server lisening on %s\n", config.Config.Server.Addr)
		if err := srv.ListenAndServe(); err != nil || err != http.ErrServerClosed {
			log.Fatalf("Listen and serve failed: %+v\n", err)
		}
	}()

	q := make(chan os.Signal, 1)
	signal.Notify(q, os.Interrupt)

	<-q
	log.Println("shutting down server...")
	var timeout time.Duration
	if config.Config.Server.ShutdownTimeout == 0 {
		timeout = ShutdownTimeout
	} else {
		timeout = time.Duration(config.Config.Server.ShutdownTimeout) * time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("could not gracefully shutdown the server: %+v\n", err)
	}
	log.Println("server shutdown completed")
}
