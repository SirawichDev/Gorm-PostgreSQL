package Api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status int
	Meta   interface{}
	Data   interface{}
}

func Res(w *gin.Context, status int, payload interface{}) {
	fmt.Print("status", status)
	var res ResponseData
	res.Status = status
	res.Data = payload
	w.JSON(200, res)
}
