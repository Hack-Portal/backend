package main

import (
	"context"
	"temp/cmd/config"
	"temp/pkg/logger"
	"temp/src/driver/db"
	"temp/src/infrastructures/server"
	"time"
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
