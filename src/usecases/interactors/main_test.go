package interactors

import (
	"fmt"
	"os"
	"testing"

	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/Hack-Portal/backend/cmd/migrations"
	gormComm "github.com/Hack-Portal/backend/src/frameworks/db/gorm"
	"github.com/golang-migrate/migrate/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db              gormComm.Connection
	dbconn          *gorm.DB
	migrateInstance *migrate.Migrate
)

func setup() {
	var err error
	// ENVを設定する
	config.LoadEnv()

	// DB接続する
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.DBName,
		config.Config.Database.Port,
		config.Config.Database.SSLMode,
		config.Config.Database.TimeZone,
	)

	dbconn, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sqlDB, _ := dbconn.DB()

	// テスト用のDBを作成する
	m, err := migrations.NewPostgresMigrate(sqlDB, "file://../../../cmd/migrations", nil)
	if err != nil {
		fmt.Println("migrate instance error:", err)
		os.Exit(1)
	}
	m.Up()
	migrateInstance = m
}

func TestMain(m *testing.M) {
	setup()
	result := m.Run()
	err := migrateInstance.Down()
	if err != nil {
		fmt.Println("migrate instance error:", err)
		os.Exit(1)
	}

	os.Exit(result)
}
