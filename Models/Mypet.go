package Models

import (
	"github.com/CRUD-Gin-Gorm/Config"
	_ "github.com/lib/pq"
)

//CreateUser of Table
func CreateUser(u *[]User) (err error) {
	err = Config.Db.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}
