package auth

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/Hack-Portal/backend/cmd/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbconn *gorm.DB

func TestMain(m *testing.M) {
	var err error
	config.LoadEnv()
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
	mig, err := migrations.NewPostgresMigrate(sqlDB, "file://../../../../cmd/migrations", nil)
	if err != nil {
		fmt.Println("migrate instance error:", err)
		os.Exit(1)
	}
	log.Println(mig.Up())
	time.Sleep(10 * time.Second)
	result := m.Run()
	if result != 0 {
		mig.Down()
		os.Exit(result)
	}

	log.Println(mig.Down())
}
