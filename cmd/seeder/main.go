package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Hack-Portal/backend/cmd/config"
	"github.com/Hack-Portal/backend/src/datastructure/models"
	"github.com/Hack-Portal/backend/src/utils/password"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	seedFile string
	user     string
)

func init() {
	flag.StringVar(&seedFile, "f", "cmd/seeder/seed.json", "path to seed file")
	flag.StringVar(&user, "u", "user", "new user")

	flag.Parse()

	if err := config.LoadEnv(); err != nil {
		log.Fatalln(err)
	}
}

type SeedFile struct {
	Role       []models.Role       `json:"role"`
	RbacPolicy []models.RbacPolicy `json:"rbacPolicies"`
}

func main() {
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
		log.Fatal("gorm open error: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("db.DB error: ", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("sqlDB ping error: ", err)
	}

	defer sqlDB.Close()

	// load Seedfile
	file, err := os.ReadFile(seedFile)
	if err != nil {
		log.Fatalf("Failed to read seed file: %v", err)
	}

	var sf SeedFile
	if err := json.Unmarshal(file, &sf); err != nil {
		log.Fatalf("Failed to unmarshal seed file: %v", err)
	}

	if err := db.Create(&sf.Role).Error; err != nil {
		log.Fatalf("Failed to create roles: %v", err)
	}

	if err := db.Create(&sf.RbacPolicy).Error; err != nil {
		log.Fatalf("Failed to create rbac policies: %v", err)
	}

	password, err := password.HashPassword(user)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	u := models.User{
		UserID:   uuid.New().String(),
		Name:     user,
		Password: password,
		Role:     1,
	}

	if err := db.Create(&u).Error; err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	fmt.Println("Seeding complete")
}
