package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/Hack-Portal/backend/cmd/migrate/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
)

var migratefile string

func init() {
	flag.StringVar(&migratefile, "f", "", "migrate file path")
	flag.Parse()
	if err := config.LoadEnv(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}

func run() error {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.DBName,
		config.Config.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("gorm open error: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("db.DB error: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("db.Ping error: %w", err)
	}

	defer sqlDB.Close()

	// migrate
	m, err := migrations.NewPostgresMigrate(sqlDB, "file://"+migratefile, nil)
	if err != nil {
		return fmt.Errorf("migrate new error: %w", err)
	}

	// migrate up
	if err := m.Up(); err != nil {
		// 変更がない場合は無視
		if err != migrate.ErrNoChange {
			return fmt.Errorf("'migrate up' error: %w", err)
		}
	}

	log.Println("'migrate up' successful")
	return nil
}
