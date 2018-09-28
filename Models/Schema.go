package Models

import (
	"github.com/jinzhu/gorm"
)

//User table
type User struct {
	gorm.Model
	FirstName string `json:",Fname"`
	LastName  string `json:",Lname"`
	PetName   string `json:",Pname"`
}

//TableName Create User struct with another name
func (u User) TableName() string {
	return "MyPet"
}
