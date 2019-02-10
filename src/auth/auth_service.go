package auth

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func FindUser(db *gorm.DB, name string, password string) (bool, LoginUser) {
	var loginUser LoginUser
	db.Table("users").Select("user_id, name, 1 as role").
		Where("name = ? AND password = ?", name, password).
		FirstOrInit(&loginUser, LoginUser{UserId: 0})
	fmt.Println("search result:" + fmt.Sprint(loginUser.UserId))
	if loginUser.UserId == 0 {
		return false, loginUser
	} else {
		return true, loginUser
	}
}
