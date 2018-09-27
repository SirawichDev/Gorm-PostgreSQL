package Models

import (
	"github.com/jinzhu/gorm"
)

//User table
type User struct {
	gorm.Model
	FirstName string `json:",omitempty"`
	LastName  string `json:",omitempty"`
	PetName   string `json:",omitempty"`
}

//TableName Create User struct with another name
func (u User) TableName() string {
	return "MyPet"
}
