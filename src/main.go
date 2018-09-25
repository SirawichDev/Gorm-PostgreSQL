package main

import (
	"github.com/CRUD-Gin-Gorm/Models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var err error

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
	User := &Models.User{}
	db.CreateTable(User)
	user := &Models.User{
		Uname:    "sirawit0676",
		Password: "1w2e3r4t5y",
		NickName: "Ex",
	}
	db.Create(user)

	println("Done!")
}
