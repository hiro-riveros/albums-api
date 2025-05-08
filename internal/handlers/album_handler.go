package handlers

import (
	"album-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	albums, err := models.GetAlbums()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Albums found"})
		return
	}

	c.JSON(http.StatusOK, albums)
}

func GetAlbumById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
	}

	album, err := models.GetAlbumById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, album)
}

func PostAlbums(c *gin.Context) {
	var newAlbum models.Album
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album, err := newAlbum.CreateAlbum()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new Album"})
		return
	}
	c.JSON(http.StatusCreated, album)
}
