package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/hackhack-Geek-vol6/backend/cmd/config"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// 必要に応じて変更する(ローカル環境なら変更する必要なし)
func setEnv(t *testing.T) {
	t.Setenv("PSQL_HOST", "localhost")
	t.Setenv("PSQL_PORT", "5432")
	t.Setenv("PSQL_USER", "root")
	t.Setenv("PSQL_PASSWORD", "postgres")
	t.Setenv("PSQL_DBNAME", "hackhack")
	t.Setenv("PSQL_SSLMODE", "disable")
	t.Setenv("PSQL_CONNECT_TIMEOUT", "10")
	t.Setenv("PSQL_CONNECT_WAIT_TIME", "10")
	t.Setenv("PSQL_CONNECT_ATTEMPTS", "3")
	t.Setenv("PSQL_CONNECT_BLOCKS", "false")
	t.Setenv("PSQL_CLOSE_TIMEOUT", "10")
}

func TestConnection(t *testing.T) {
	// setEnv(t)
	// defer os.Clearenv()
	config.LoadEnv()

	conn := NewConnection()

	db, err := conn.Connection()
	if err != nil {
		t.Errorf("failed to connect to database :%v", err)
	}

	t.Log("success to connect database")

	sql, err := db.DB()
	if sql.Ping() != nil {
		t.Errorf("failed to ping database :%v", err)
	}

	t.Log("success to ping database")

	if conn.Close(context.Background()) != nil {
		t.Errorf("failed to close database :%v", err)
	}

	t.Log("success to close database")
}
