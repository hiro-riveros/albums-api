package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	for i, album := range albums {
		fmt.Printf("i: %v", i)

		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  string `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "Jhon Coltrane", Price: "56.99"},
	{ID: "2", Title: "Insomniac", Artist: "Green Day", Price: "39.99"},
	{ID: "3", Title: "American Idiot", Artist: "Green Day", Price: "49.99"},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
