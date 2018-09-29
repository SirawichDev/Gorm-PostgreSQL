package main

import (
	"github.com/CRUD-Gin-Gorm/Route"

	"github.com/CRUD-Gin-Gorm/Config"
	"github.com/CRUD-Gin-Gorm/Models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var err error

//Data sample
var users = []Models.User{
	Models.User{FirstName: "Sirawih", LastName: "Voungchuy", PetName: "Garfiled"},
	Models.User{FirstName: "Samon", LastName: "Miew", PetName: "jimmy"},
	Models.User{FirstName: "Kann", LastName: "naruuu", PetName: "kanuu"},
	Models.User{FirstName: "Kob", LastName: "boob", PetName: "Bonny"},
}

func main() {
	Config.Db, err = gorm.Open("postgres", "user=sirawich dbname=booker sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer Config.Db.Close()

	dbase := Config.Db.DB()
	defer dbase.Close()

	err = dbase.Ping()
	if err != nil {
		panic(err.Error())
	}
	println("Connection to database established")
	// Config.Db.CreateTable(&Models.User{})

	// u := []Models.User{}

	//update
	// u := Models.User{FirstName: "Kann"}
	// Models.Update(&u)
	// db.Where(&u).First(&u)
	// fmt.Println(u)
	// u.Uname = "Sirawich"
	// db.Save(&u)

	//Delete
	// Models.Delete(&u)
	// db.Where(&Models.User{Uname: "Huzen"}).Delete(&Models.User{})

	//show first data in table
	// Models.GetOne(&u)

	//show last data int table
	// Models.GetOneLast(&u)

	// fmt.Println(u)
	println("Done!")
	r := Route.InitRouter()
	r.Run()

}
