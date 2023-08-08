package repository_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	bootstrap "github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"
	_ "github.com/lib/pq"
)

var testQueries *repository.Queries

func TestMain(m *testing.M) {
	config := bootstrap.LoadEnvConfig("../../../../../../")
	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannnot connection DB :", err)
	}

	testQueries = repository.New(testDB)

	os.Exit(m.Run())
}
