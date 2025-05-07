package routes

import (
	"album-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAlbumRoutes(router *gin.Engine, albumHandler *handlers.AlbumHandler) {
	router.GET("/albums", albumHandler.GetAlbums)
	router.GET("/albums/:id", albumHandler.GetAlbumById)
	router.POST("/albums", albumHandler.PostAlbums)
}
