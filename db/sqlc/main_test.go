package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/hackhack-Geek-v6/backend/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadEnvConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSouse)
	if err != nil {
		log.Fatal("cannnot open sql:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
