package main

import (
	"fmt"

	"github.com/CRUD-Gin-Gorm/Models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var err error

//Data sample
var users []Models.User = []Models.User{
	Models.User{Uname: "x1", Password: "1234", NickName: "xx"},
	Models.User{Uname: "x2", Password: "3241", NickName: "xxx"},
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
	db.CreateTable(&Models.User{})
	for _, users := range users {
		db.Create(&users)
	}
	// user := &Models.User{
	// 	Uname:    "sirawit0676",
	// 	Password: "1w2e3r4t5y",
	// 	NickName: "Ex",
	// }
	u := Models.User{}
	//show first data in table
	// db.First(&u)
	//show last data intable
	db.Last(&u)
	fmt.Println(u)
	println("Done!")
}
