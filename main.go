package main

import (
	"database/sql"
	"log"

	"github.com/hackhack-Geek-vol6/backend/api"
	db "github.com/hackhack-Geek-vol6/backend/db/sqlc"
	_ "github.com/hackhack-Geek-vol6/backend/docs"
	"github.com/hackhack-Geek-vol6/backend/util"
	_ "github.com/lib/pq"
)

//	@title			Geek Hackathon vol6 backend API
//	@version		1.0
//	@description	Geek Camp vol6で作ったAPI
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/v1
func main() {
	config, err := util.LoadEnvConfig(".")
	if err != nil {
		log.Fatal("cannnot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSouse)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal(err)
	}
	if err := server.Start("127.0.0.1:8080"); err != nil {
		log.Fatal("cannnot start server :", err)
	}
}
