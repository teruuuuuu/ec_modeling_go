package filter

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"../auth"
	"../session"
)

type LoginForm struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func loginCheck(c *gin.Context, as *session.AppSession, db *gorm.DB) {
	if c.Request.URL.Path == "/auth/login" && c.Request.Method == "POST" {
		isLogin, _ := as.GetLoginUser()
		if isLogin {
			c.JSON(400, gin.H{"result": "fail", "message": "already logined"})
			c.Abort() // 処理続行停止
			return
		}
		loginForm := LoginForm{}
		c.BindJSON(&loginForm)
		isValid, findUser := auth.FindUser(db, loginForm.Name, loginForm.Password)
		if isValid {
			as.SetLoginUser(&auth.LoginUser{UserId: findUser.UserId, Name: findUser.Name, Role: findUser.Role})
			c.JSON(200, gin.H{"result": "success"})
			c.Abort() // 処理続行停止
		} else {
			c.JSON(400, gin.H{"result": "fail"})
			c.Abort() // 処理続行停止
		}
	} else if c.Request.URL.Path == "/auth/logout" && c.Request.Method == "POST" {
		as.Clear()
		c.JSON(200, gin.H{"result": "success"})
		c.Abort() // 処理続行停止
	} else {
		isLogin, loginUser := as.GetLoginUser()
		fmt.Println(loginUser)
		if !isLogin {
			c.JSON(400, gin.H{"error": "no_login"})
			c.Abort() // 処理続行停止
		}
	}
}
