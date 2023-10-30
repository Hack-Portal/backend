package main

import (
	"context"
	"time"

	"github.com/hackhack-Geek-vol6/backend/cmd/config"
	"github.com/hackhack-Geek-vol6/backend/pkg/logger"
	"github.com/hackhack-Geek-vol6/backend/src/driver/db"
	"github.com/hackhack-Geek-vol6/backend/src/infrastructures/server"
)

func main() {
	l := logger.NewLogger(logger.DEBUG)
	config.LoadEnv(l)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Config.Cockroach.CloseTimeout)*time.Second)
	defer cancel()

	conn := db.NewConnection(l)
	defer conn.Close(ctx)

	db, err := conn.Connection()
	if err != nil {
		l.Errorf("database connecting err :%v", err)
		return
	}

	server.NewServer(db, l).Run()
}
