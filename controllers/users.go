package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masiuciszek/go-server/models"
	"github.com/masiuciszek/go-server/utils"
)

// UserInput fields required to create ne user
type UserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

// UpdateUserInput input for updating user
type UpdateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      string `json:"age"`
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
	var user models.User

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	hashedPassword, err := utils.HashPassword(userInput.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := models.User{Name: userInput.Name, Email: userInput.Email, Password: hashedPassword, Age: userInput.Age}
	models.DB.Create(&newUser)

	// name, value string, maxAge int, path, domain string, secure, httpOnly bool
	c.SetCookie("auth_cookie", token, 36000, "/", "localhost", false, false) // this should be in a ENV file

	c.JSON(http.StatusOK, gin.H{"data": newUser, "token": token})
}

// UpdateUser route
// PUT user/:id
func UpdateUser(c *gin.Context) {

	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&user).Update(input)

	c.JSON(http.StatusOK, gin.H{"data": user})

}

// GetUserByID route
// GET user/:id
func GetUserByID(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
