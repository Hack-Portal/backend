package main

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/src/drivers/postgres"
	"github.com/hackhack-Geek-vol6/backend/src/infrastructure/server"
)

func main() {
	conn := postgres.NewConnection()
	defer conn.Close(context.Background())

	dbconn, err := conn.Connection()
	if err != nil {
		panic(err)
	}

	server.NewGinServer(dbconn).Run()

}
