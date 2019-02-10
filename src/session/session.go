package session

import (
	"../auth"
	"github.com/gin-contrib/sessions"
)

type AppSession struct {
	S *sessions.Session
}

func (appSession *AppSession) Clear() {
	(*appSession.S).Clear()
	(*appSession.S).Save()
}

func (appSession *AppSession) GetCount() (bool, int) {
	v := (*appSession.S).Get("count")
	if v == nil {
		return false, 0
	} else {
		return true, v.(int)
	}
}

func (appSession *AppSession) SetCount(count int) {
	(*appSession.S).Set("count", count)
	(*appSession.S).Save()
}

func (appSession *AppSession) GetLoginUser() (bool, *auth.LoginUser) {
	id := (*appSession.S).Get("login_user_id")
	name := (*appSession.S).Get("login_user_name")
	role := (*appSession.S).Get("login_user_role")
	if id == nil {
		return false, &auth.LoginUser{}
	} else {
		return true, &auth.LoginUser{UserId: id.(uint), Name: name.(string), Role: role.(string)}
	}
}

func (appSession *AppSession) SetLoginUser(loginUser *auth.LoginUser) {
	(*appSession.S).Set("login_user_id", loginUser.UserId)
	(*appSession.S).Set("login_user_name", loginUser.Name)
	(*appSession.S).Set("login_user_role", loginUser.Role)
	(*appSession.S).Save()
}
