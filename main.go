package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/masiuciszek/go-server/controllers"
	"github.com/masiuciszek/go-server/models"
)

func main() {
	r := gin.Default()

	models.ConnectDb()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "Legia warszawa"})
	// })
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowHeaders: []string{"Origin"},
	}))

	r.GET("/users", controllers.GetUsers)

	r.GET("/dishes", controllers.GetDishes)
	r.POST("/dish", controllers.CreateDish)
	r.GET("/dish/:id", controllers.GetDishByID)
	r.PUT("/dish/:id", controllers.UpdateDish)
	r.DELETE("/dish/:id", controllers.RemoveDish)

	r.Run()
}
