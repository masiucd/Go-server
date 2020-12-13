package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masiuciszek/go-server/models"
)

// UserInput fields required to create ne user
type UserInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Age   int    `json:"age" binding:"required"`
}

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

// RegisterUser route
// POST /user
func RegisterUser(c *gin.Context) {
	var userInput UserInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := models.User{Name: userInput.Name, Email: userInput.Email, Age: userInput.Age}
	models.DB.Create(&newUser)

	c.JSON(http.StatusOK, gin.H{"data": newUser})
}
