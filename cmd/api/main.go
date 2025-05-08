package main

import (
	"album-api/internal/handlers"
	"album-api/internal/middlewares"
	"album-api/internal/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()

	// Gin Init
	router := gin.Default()
	publicRoutes := router.Group("/")
	publicRoutes.POST("/login", handlers.Login)
	publicRoutes.POST("/registration", handlers.Register)
	// Router
	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middlewares.JwtAuthMiddleware())
	protectedRoutes.GET("/albums", handlers.GetAlbums)
	protectedRoutes.GET("/albums/:id", handlers.GetAlbumById)
	protectedRoutes.POST("/albums", handlers.PostAlbums)

	// Start Server
	port := ":8080"
	fmt.Printf("Server listening on port %s\n", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
