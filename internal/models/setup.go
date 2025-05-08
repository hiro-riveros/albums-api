package models

import (
	"album-api/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDataBase() {
	dbConfig, dbError := config.LoadDBConfig()
	if dbError != nil {
		log.Fatalf("Error loading config: %v", dbError)
	}

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.SSLMode)
	db, dbError = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if dbError != nil {
		log.Fatalf("Failed to connect to database: %v", dbError)
	}
	fmt.Println("Successfully connected to PostgreSQL database!")

	// Auto Migrate
	db.AutoMigrate(&Album{}, &User{})
}
