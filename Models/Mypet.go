package Models

import (
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
	err = Config.Db.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

//GetOneFirst Get data At first of record
func GetOneFirst(u *[]User) (err error) {
	err = Config.Db.Where(&u).First(&u).Error
	if err != nil {
		return err
	}
	return nil
}

//GetOneLast Get data At last of record
func GetOneLast(u *[]User) (err error) {
	err = Config.Db.Where(&u).Last(&u).Error
	if err != nil {
		return err
	}
	return nil
}

//Delete Data
func Delete(u *User) (err error) {
	err = Config.Db.Where(&u).Delete(&u).Error
	if err != nil {
		return err
	}
	return nil
}

//Update Data
func Update(u *User) (err error) {
	err = Config.Db.Where(&u).First(&u).Error
	if err != nil {
		return err
	}
	u.PetName = "monmon"
	Config.Db.Save(&u)
	return nil
}
