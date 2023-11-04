package main

import (
	"context"

	"github.com/hackhack-Geek-vol6/backend/src/drivers/postgres"
	"github.com/hackhack-Geek-vol6/backend/src/infrastructure/server"
)

//	@title						Hack Hack Backend API
//	@version					0.1.0
//	@description			Hack Hack Backend API serice
//	@termsOfService	　https://api.seafood-dev.com

//	@contact.name			murasame29
//	@contact.url			https://twitter.com/fresh_salmon256
//	@contact.email		oogiriminister@gmail.com

//	@license.name			No-license
//	@license.url			No-license

// @host							https://api.seafood-dev.com
// @BasePath					/v1
func main() {
	conn := postgres.NewConnection()
	defer conn.Close(context.Background())

	dbconn, err := conn.Connection()
	if err != nil {
		panic(err)
	}

	server.NewGinServer(dbconn).Run()

}
