package db

import (
	"database/sql"
	"fmt"

	"github.com/Hack-Portal/backend/cmd/config"
	_ "github.com/lib/pq"
	"go.nhat.io/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDBWithOtelSQL はOtelでwrapされたDBを返す
func ConnectDBWithOtelSQL() *sql.DB {
	driverName, err := otelsql.Register(
		config.Config.Database.Driver,
		otelsql.AllowRoot(),
		otelsql.TraceQueryWithArgs(),
		otelsql.TraceRowsClose(),
		otelsql.TraceRowsAffected(),
		otelsql.WithDatabaseName(config.Config.Database.DBName),
		otelsql.WithSystem(semconv.DBSystemPostgreSQL),
	)
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.DBName,
		config.Config.Database.SSLMode,
	)

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}

	return db
}

// ConnectDB DBのを返す
func ConnectDB() *sql.DB {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.DBName,
		config.Config.Database.SSLMode,
	)

	db, err := sql.Open(config.Config.Database.Driver, dsn)
	if err != nil {
		panic(err)
	}

	// TODO: Connectionのリトライ処理を追加する
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}

// NewGORM は*sql.DBからGORMのifを返す
func NewGORM(db *sql.DB) *gorm.DB {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))

	if err != nil {
		panic(err)
	}

	return gormDB
}
