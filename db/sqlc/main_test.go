package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/hackhack-Geek-vol6/backend/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadEnvConfig("../../")
	if err != nil {
		log.Fatal("cannnot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSouse)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
