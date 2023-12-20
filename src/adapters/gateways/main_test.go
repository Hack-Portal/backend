package gateways

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/Hack-Portal/backend/cmd/migrations"
	"github.com/Hack-Portal/backend/src/driver/aws"
	gormComm "github.com/Hack-Portal/backend/src/frameworks/db/gorm"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/golang-migrate/migrate/v4"
	"github.com/murasame29/db-conn/sqldb/postgres"
	"gorm.io/gorm"
)

var (
	db              gormComm.Connection
	dbconn          *gorm.DB
	migrateInstance *migrate.Migrate
	client          *s3.Client
)

func setup() {
	var err error
	// ENVを設定する
	config.LoadEnv()

	// DB接続する
	postgresConn := postgres.NewConnection(
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.DBName,
		config.Config.Database.SSLMode,
		config.Config.Database.ConnectAttempts,
		config.Config.Database.ConnectWaitTime,
		config.Config.Database.ConnectBlocks,
	)

	sqlDB, err := postgresConn.Connection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db = gormComm.New()
	dbconn, err = db.Connection(sqlDB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// テスト用のDBを作成する
	m, err := migrations.NewPostgresMigrate(sqlDB, "file://../../../cmd/migrations", nil)
	if err != nil {
		fmt.Println("migrate instance error:", err)
		os.Exit(1)
	}
	m.Up()

	migrateInstance = m

	// AWS S3に接続する

	client, err = aws.New().Connect(context.TODO())
	if err != nil {
		fmt.Println("aws connection error :", err)
		os.Exit(1)
	}
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
