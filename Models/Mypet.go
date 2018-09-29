package Models

import (
	"fmt"

	"github.com/CRUD-Gin-Gorm/Config"
	_ "github.com/lib/pq"
)

//single create
// user := &Models.User{
// 	Uname:    "sirawit0676",
// 	Password: "1w2e3r4t5y",
// 	NickName: "Ex",
// }

//CreateUser of Table
func CreateUser(u *User) (err error) {
	err = Config.Db.Create(u).Error
	if err != nil {
		return err
	}
	return nil
}

//GetAllUser
func GetAllUser(u *[]User) (err error) {
	err = Config.Db.Find(u).Error
	if err != nil {
		return err
	}
	return nil
}

//GetOne
func GetOne(u *User, id string) (err error) {
	err = Config.Db.Where(id).First(u).Error
	if err != nil {
		return err
	}
	return nil
}

//GetOneFirst Get data At first of record
func GetOneFirst(u *[]User) (err error) {
	err = Config.Db.Where(u).First(u).Error
	if err != nil {
		return err
	}
	return nil
}

//GetOneLast Get data At last of record
func GetOneLast(u *[]User) (err error) {
	err = Config.Db.Where(u).Last(u).Error
	if err != nil {
		return err
	}
	return nil
}

//Delete Data
func Delete(u *User, id string) (err error) {
	Config.Db.Where(id).Delete(u)
	return nil
}

//Update Data
func Update(u *User, id string) (err error) {
	fmt.Println(u)
	Config.Db.Save(&u)
	return nil
}
