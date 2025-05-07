package main

import (
	"album-api/config"
	"album-api/internal/handlers"
	"album-api/internal/models"
	"album-api/internal/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)
	db, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Successfully connected to PostgreSQL database!")

	// Auto Migrate
	db.AutoMigrate(&models.Album{})

	// Gin Init
	router := gin.Default()

	// Model Instance
	albumHander := handlers.NewAlbumHandler(db)

	// Router
	routes.SetupAlbumRoutes(router, albumHander)

	// Start Server
	port := ":8080"
	fmt.Printf("Server listening on port %s\n", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
