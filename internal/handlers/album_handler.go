package handlers

import (
	"album-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AlbumHandler struct {
	db *gorm.DB
}

func NewAlbumHandler(db *gorm.DB) *AlbumHandler {
	return &AlbumHandler{db: db}
}

func (h *AlbumHandler) GetAlbums(c *gin.Context) {
	var albums []models.Album
	result := h.db.Find(&albums)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch albums"})
		return
	}

	c.JSON(http.StatusOK, albums)
}

func (h *AlbumHandler) GetAlbumById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
	}

	var album models.Album
	result := h.db.First(&album, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch album"})
		return
	}
	c.JSON(http.StatusOK, album)
}

func (h *AlbumHandler) PostAlbums(c *gin.Context) {
	var newAlbum models.Album
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.db.Create(&newAlbum)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create new Album"})
		return
	}
	c.JSON(http.StatusCreated, newAlbum)
}
