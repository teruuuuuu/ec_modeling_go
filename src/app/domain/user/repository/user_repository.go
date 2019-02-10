package user_repository

import (
	"github.com/jinzhu/gorm"

	user_aggregate "../aggregate"
	user_model "../model"
)

func DeleteAll(db *gorm.DB) {
	db.Delete(user_model.User{})
	db.Delete(user_model.UserInfo{})
}

func Save(db *gorm.DB, ua *user_aggregate.UserAggregate) {
	db.Save(&ua.User)
	ua.UserInfo.UserId = &ua.User.UserId
	db.Save(&ua.UserInfo)
}
