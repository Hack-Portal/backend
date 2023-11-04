package postgres

import (
	"context"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestConnection(t *testing.T) {
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
