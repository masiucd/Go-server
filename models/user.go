package models

// User type model
type User struct {
	
	ID     uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
}