package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hackhack-Geek-vol6/backend/cmd/config"
)

func RunWithGracefulShutdown(router interface{}) {
	var srv http.Server
	switch v := router.(type) {
	case *gin.Engine:
		srv = http.Server{
			Addr:    config.Config.Server.Addr,
			Handler: v,
		}
	default:
		panic("router is not gin.Engine")
	}

	go func() {
		fmt.Println("server starting...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// graceful shutdown
	q := make(chan os.Signal)
	signal.Notify(q, os.Interrupt, os.Kill)

	<-q
	fmt.Println("server stopping...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}
	fmt.Println("server stopped")
}
