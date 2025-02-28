package database

import (
	"fmt"
	"log"

	"github.com/CodeATM/notepal-go/config"
	"github.com/CodeATM/notepal-go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb(cfg config.Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.PostgresServer,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
		cfg.PostgresPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to the database successfully")

	// Add UUID extension
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// Migrate models
	DB.AutoMigrate(models.Models()...)
	log.Println("Database migrations completed")
}
