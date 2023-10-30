package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"temp/cmd/config"
	"time"

	"github.com/gin-gonic/gin"
)

func RunWithGracefulStop(router any) {
	var srv *http.Server

	switch v := router.(type) {
	case *gin.Engine:
		srv = &http.Server{Addr: config.Config.Server.Addr, Handler: v}
	default:
		log.Fatal("Invalid router type")
		return
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	s := make(chan os.Signal)
	signal.Notify(s, os.Interrupt, os.Kill)

	<-s
	log.Println("server shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.Server.ShutdownTimeout)*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown:", err)
	}

	log.Println("server exiting")
}
