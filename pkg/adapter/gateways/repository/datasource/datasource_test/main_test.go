package repository_test

import (
	"database/sql"
	"os"
	"testing"

	repository "github.com/hackhack-Geek-vol6/backend/pkg/adapter/gateways/repository/datasource"
	bootstrap "github.com/hackhack-Geek-vol6/backend/pkg/bootstrap"

	// util "github.com/hackhack-Geek-vol6/backend/pkg/util/password"
	_ "github.com/lib/pq"
)

var testQueries *repository.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	bootstrap.LoadEnvConfig("../../../../../../")

	testQueries = repository.New(testDB)

	os.Exit(m.Run())
}
