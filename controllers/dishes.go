package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masiuciszek/go-server/models"
)

// CreateDishInput input
type CreateDishInput struct {
	Name string `json:"name" binding:"required"`
	Stars int  `json:"stars" binding:"required"`	
  Author string `json:"author" binding:"required"`
}

// GetDishes get all dishes from db
// GET /books
func GetDishes(c *gin.Context) {
	var dishes []models.Dish
	models.DB.Find(&dishes)

	c.JSON(http.StatusOK, gin.H{"data": dishes})
}


// CreateDish route
// POST /dish
func CreateDish(c *gin.Context)  {
	var input CreateDishInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
	}

	dish := models.Dish{Name: input.Name, Stars: input.Stars, Author: input.Author }
	models.DB.Create(dish)

	c.JSON(http.StatusOK, gin.H{"data":dish})
}