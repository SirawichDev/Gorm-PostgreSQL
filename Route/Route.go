package Route

import (
	"github.com/CRUD-Gin-Gorm/Controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("User", Controllers.AllUser)
		v1.POST("User", Controllers.CreateUser)
	}
	return r
}
