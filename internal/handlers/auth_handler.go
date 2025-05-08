package handlers

import (
	"album-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var anonUser models.User

	if err := c.ShouldBindJSON(&anonUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := models.Login(anonUser.Email, anonUser.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, token)
}

func Register(c *gin.Context) {
	var anonUser models.User

	if err := c.ShouldBindJSON(&anonUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:    anonUser.Email,
		Password: anonUser.Password,
	}
	_, err := user.SaveUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration OK", "user": user})
}
