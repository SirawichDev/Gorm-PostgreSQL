package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var err error

type User struct {
	ID        uint
	Username  string
	FirstName string
	LastName  string
}

func main() {
	db, err := gorm.Open("postgres", "user=sirawich dbname=booker sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dbase := db.DB()
	defer dbase.Close()

	err = dbase.Ping()
	if err != nil {
		panic(err.Error())
	}
	println("Connection to database established")
	db.CreateTable(&User{})
	user := User{
		Username:  "sirawit0676",
		FirstName: "Sirawich",
		LastName:  "Vougnchuy",
	}
	db.Create(&user)

	println("Done!")
}
