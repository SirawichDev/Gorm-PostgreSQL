package Controllers

import (
	"github.com/CRUD-Gin-Gorm/Api"
	"github.com/CRUD-Gin-Gorm/Models"
	"github.com/gin-gonic/gin"
)

func AllUser(c *gin.Context) {
	u := []Models.User{}
	err := Models.GetAllUser(&u)
	if err != nil {
		Api.Res(c, 200, &u)
	}

}
