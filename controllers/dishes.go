package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masiuciszek/go-server/models"
)

// CreateDishInput input
type CreateDishInput struct {
	Name   string `json:"name" binding:"required"`
	Stars  int    `json:"stars" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// UpdateDishInput input for updating dish
type UpdateDishInput struct {
	Name   string `json:"name"`
	Author string `json:"author"`
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
func CreateDish(c *gin.Context) {
	var input CreateDishInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dish := models.Dish{Name: input.Name, Stars: input.Stars, Author: input.Author}
	models.DB.Create(&dish)

	c.JSON(http.StatusOK, gin.H{"data": dish})
}

// GetDishByID route
// GET /dish/:id
func GetDishByID(c *gin.Context) {
	var dish models.Dish

	if err := models.DB.Where("id = ?", c.Param("id")).First(&dish).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dish})
}

// UpdateDish route
// PUT /dish/:id
func UpdateDish(c *gin.Context) {
	var dish models.Dish

	if err := models.DB.Where("id = ?", c.Param("id")).First(&dish).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	// validate input
	var input UpdateDishInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&dish).Update(input)

	c.JSON(http.StatusOK, gin.H{"data": dish})

}

// RemoveDish route
// DELETE /dish/:id
func RemoveDish(c *gin.Context) {
	var dish models.Dish

	if err := models.DB.Where("id = ?", c.Param("id")).First(&dish).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	models.DB.Delete(&dish)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
