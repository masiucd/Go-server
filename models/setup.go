package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB init
var DB *gorm.DB

// ConnectDb func
func ConnectDb()  {
	database, err := gorm.Open("sqlite3", "test.db")

  if err != nil {
    panic("Failed to connect to database!")
  }

  database.AutoMigrate(&Dish{})

  DB = database
}