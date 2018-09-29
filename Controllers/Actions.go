package Controllers

import (
	"fmt"

	"github.com/CRUD-Gin-Gorm/Api"
	"github.com/CRUD-Gin-Gorm/Models"
	"github.com/gin-gonic/gin"
)

func AllUser(c *gin.Context) {
	u := []Models.User{}
	err := Models.GetAllUser(&u)
	fmt.Println(u)
	if err != nil {
		Api.Res(c, 404, u)
	} else {
		Api.Res(c, 200, u)
	}

}
func CreateUser(c *gin.Context) {
	user := Models.User{}
	user.FirstName = c.PostForm("fname")
	user.LastName = c.PostForm("lname")
	user.PetName = c.PostForm("pname")
	c.BindJSON(&user)
	fmt.Printf("%v", user)
	err := Models.CreateUser(&user)
	if err != nil {
		Api.Res(c, 404, user)
	} else {
		Api.Res(c, 200, user)
	}

}
