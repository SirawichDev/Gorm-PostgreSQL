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
	Models.User{Uname: "Sirawih", Password: "1234", NickName: "Ex"},
	Models.User{Uname: "Huzen", Password: "3241", NickName: "Sainee"},
	Models.User{Uname: "Punupong", Password: "1w2e3r", NickName: "Tar"},
	Models.User{Uname: "Chakrit", Password: "3e4r5y", NickName: "Him"},
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
	// u := Models.User{}

	//update
	u := Models.User{NickName: "Ex"}
	db.Where(&u).First(&u)
	fmt.Println(u)

	u.Uname = "Sirawich"
	db.Save(&u)
	//show first data in table
	// db.First(&u)
	//show last data int table
	// db.Last(&u)
	fmt.Println(u)
	println("Done!")

}
