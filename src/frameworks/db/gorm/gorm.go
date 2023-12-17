package gorm

import (
	"context"
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection interface {
	Connection(*sql.DB) (*gorm.DB, error)
	Close(ctx context.Context) error
}

type connection struct {
	db *gorm.DB
}

func New() Connection {
	return &connection{}
}

func (c *connection) Connection(db *sql.DB) (*gorm.DB, error) {
	var err error
	c.db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))
	if err != nil {
		return nil, err
	}

	return c.db, nil
}

func (c *connection) Close(ctx context.Context) error {
	sqlDB, err := c.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
