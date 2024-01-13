package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/bwmarrin/discordgo"
)

type httpServer struct {
	srv *http.Server
}

func New() *httpServer {
	return &httpServer{}
}

func (s *httpServer) Run(handler http.Handler) {
	s.run(handler)
}

func (s *httpServer) RunDiscordBot(sess *discordgo.Session) {
	if err := sess.Open(); err != nil {
		panic(err)
	}
}

func (s *httpServer) CloseDiscordBot(sess *discordgo.Session) {
	if err := sess.Close(); err != nil {
		panic(err)
	}
}

func (s *httpServer) GracefulShutdown() {
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

	if err := s.srv.Shutdown(ctx); err != nil {
		log.Fatalf("could not gracefully shutdown the server: %+v\n", err)
	}
	log.Println("server shutdown completed")
}
