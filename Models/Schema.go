package Models

import (
	"github.com/jinzhu/gorm"
)

//User table
type User struct {
	gorm.Model
	FirstName string `json:"fname" form:"fname" binding:"required"`
	LastName  string `json:"lname" form:"lname" binding:"required"`
	PetName   string `json:"pname" form:"pname" binding:"required"`
}

//TableName Create User struct with another name
func (u User) TableName() string {
	return "MyPet"
}
