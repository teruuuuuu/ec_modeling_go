package user_aggregate

import user_model "../model"

type UserAggregate struct {
	User     user_model.User
	UserInfo user_model.UserInfo
}
