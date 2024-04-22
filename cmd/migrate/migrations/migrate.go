package migrations

import (
	"database/sql"
	"fmt"

	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/cockroachdb"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func NewPostgresMigrate(db *sql.DB, file string, arg any) (*migrate.Migrate, error) {
	switch config.Config.Database.Driver {
	case "postgres":
		if arg == nil {
			arg = &postgres.Config{}
		}
		return migratePostgres(db, file, arg.(*postgres.Config))
	case "cockroachdb":
		if arg == nil {
			arg = &cockroachdb.Config{}
		}
		return migrateCockroachdb(db, file, arg.(*cockroachdb.Config))
	default:
		return nil, fmt.Errorf("invalid database driver: %s", config.Config.Database.DBName)
	}
}

func migrateCockroachdb(db *sql.DB, file string, arg *cockroachdb.Config) (*migrate.Migrate, error) {
	driver, err := cockroachdb.WithInstance(db, arg)
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(
		file,
		"migrate",
		driver,
	)
}

func migratePostgres(db *sql.DB, file string, arg *postgres.Config) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(db, arg)
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(
		file,
		"postgres",
		driver,
	)
}
