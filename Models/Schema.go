package Models

import (
	"github.com/jinzhu/gorm"
)

//User table
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	PetName   string
}

//TableName Create User struct with another name
func (u User) TableName() string {
	return "MyPet"
}
