package migrations

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type migration struct {
	db *migrate.Migrate
}

func NewPostgresMigrate(db *sql.DB, file string, arg *postgres.Config) (*migration, error) {
	driver, err := postgres.WithInstance(db, arg)
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		file,
		"postgres",
		driver,
	)

	return &migration{
		db: m,
	}, nil
}

func (m *migration) Up() error {
	return m.db.Up()
}

func (m *migration) Down() error {
	return m.db.Down()
}
