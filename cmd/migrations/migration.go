package migrations

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func NewPostgresMigrate(db *sql.DB, file string, arg *postgres.Config) (*migrate.Migrate, error) {
	if arg == nil {
		arg = &postgres.Config{}
	}
	driver, err := postgres.WithInstance(db, arg)
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		file,
		"postgres",
		driver,
	)

	return m, err
}
