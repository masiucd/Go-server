package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masiuciszek/go-server/models"
)

// GetUsers route
// GET /users
func GetUsers(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)
	if len(users) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": "no users in db"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}
