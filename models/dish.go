package models

// Dish type for our Dish model
type Dish struct {
	ID uint `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Stars int `json:"stars"`
	Author string `json:"author"`
}

