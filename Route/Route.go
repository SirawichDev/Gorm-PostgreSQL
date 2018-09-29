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
		v1.POST("Sample", Controllers.SampleUser)
		v1.GET("User/:id", Controllers.GetOneUser)
		v1.PUT("User/:id", Controllers.UpdateUser)
	}
	return r
}
