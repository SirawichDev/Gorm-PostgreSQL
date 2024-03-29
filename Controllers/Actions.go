package Controllers

import (
	"fmt"

	"github.com/CRUD-Gin-Gorm/Api"
	"github.com/CRUD-Gin-Gorm/Models"
	"github.com/gin-gonic/gin"
)

//AllUser Show All User in database
func AllUser(c *gin.Context) {
	var users []Models.User
	err := Models.GetAllUser(&users)
	fmt.Println(users)
	if err != nil {
		Api.Res(c, 404, users)
	} else {
		Api.Res(c, 200, users)
	}

}
func CreateUser(c *gin.Context) {
	user := Models.User{}
	// user.FirstName = c.PostForm("fname")
	// user.LastName = c.PostForm("lname")
	// user.PetName = c.PostForm("pname")
	c.BindJSON(&user)
	fmt.Printf("%v", user)
	err := Models.CreateUser(&user)
	if err != nil {
		Api.Res(c, 404, user)
	} else {
		Api.Res(c, 200, user)
	}
}

//SampleUser make sample data to database
func SampleUser(c *gin.Context) {
	var users []Models.User = []Models.User{
		Models.User{FirstName: "sorawat", LastName: "yo", PetName: "ramzy"},
		Models.User{FirstName: "salamut", LastName: "jo", PetName: "fozy"},
		Models.User{FirstName: "muzu", LastName: "ko", PetName: "loly"},
		Models.User{FirstName: "zuyu", LastName: "lo", PetName: "maly"},
	}
	c.BindJSON(&users)
	for _, users := range users {
		err := Models.CreateUser(&users)
		if err != nil {
			Api.Res(c, 404, users)
		} else {
			Api.Res(c, 404, users)
		}
	}
}

//GetOneUser get user data
func GetOneUser(c *gin.Context) {
	var users Models.User
	id := c.Params.ByName("id")
	err := Models.GetOne(&users, id)
	fmt.Println(users)
	if err != nil {
		Api.Res(c, 404, users)
	} else {
		Api.Res(c, 200, users)
	}
}

//Update User data
func UpdateUser(c *gin.Context) {
	users := Models.User{}
	id := c.Params.ByName("id")
	err := Models.GetOne(&users, id)
	fmt.Println(users.FirstName)
	if err != nil {
		Api.Res(c, 404, users)
	}
	c.BindJSON(&users)
	//fmt.Println("From JSON to Struct", reflect.TypeOf(users))
	err = Models.Update(&users, id)
	if err != nil {
		Api.Res(c, 404, users)
	} else {
		Api.Res(c, 200, users)
	}
}

//Delete data from database
func Deleter(c *gin.Context) {
	var users Models.User
	id := c.Params.ByName("id")
	err := Models.Delete(&users, id)
	if err != nil {
		Api.Res(c, 404, users)
	} else {
		Api.Res(c, 200, users)
	}
}
