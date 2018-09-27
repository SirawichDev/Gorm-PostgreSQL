package Models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Uname    string
	Password string
	NickName string
}
