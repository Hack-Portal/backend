package repository_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	bootstrap "github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"

	// util "github.com/hackhack-Geek-vol6/backend/pkg/util/password"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := bootstrap.LoadEnvConfig("../../../../../../")
	if err != nil {
		log.Fatal("cannnot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
