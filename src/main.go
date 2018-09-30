package main

import (
	"github.com/CRUD-Gin-Gorm/Route"

	"github.com/CRUD-Gin-Gorm/Config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var err error

func main() {
	Config.Db, err = gorm.Open("postgres", "user=sirawich dbname=booker sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer Config.Db.Close()

	dbase := Config.Db.DB()
	err = dbase.Ping()
	if err != nil {
		panic(err.Error())
	}
	println("Connection to database established")
	println("Done!")
	r := Route.InitRouter()
	r.Run()

}
