package model

import (
	"github.com/jinzhu/gorm"
	"version1/utils/errmsg"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" form:"username"`
	Password string `gorm:"type:varchar(20);not null " json:"password" form:"password"`
	Email    string `gorm:"type:varchar(20);not null" json:"email" form:"email"`
}

type ParamUser struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func CheckUser(email string) int {
	var user User
	db.Select("id").Where("email = ?", email).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USEREMAIL_USED
	}
	return errmsg.SUCCESS
}

func Login(u *ParamUser) int {
	var user User
	db.Where("email = ?", u.Email).First(&user)

	if u.Password != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	return errmsg.SUCCESS
}

func Register(u *User) int {

	//写入数据库
	err := db.Create(&u).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS

}
