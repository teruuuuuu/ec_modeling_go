package main

import (
	"./app"
	"./controller"
	"./filter"
	"./infla/dao"
	"./session"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {
	var db = dao.GetDb()
	defer db.Close() // 最後にクローズする
	app := app.New(db)
	var engine = initEngine(app, db)

	engine.Run(":8000")
}

func initEngine(app *app.App, db *gorm.DB) *gin.Engine {
	var engine = gin.Default()
	store := cookie.NewStore([]byte("secret"))
	// engine.Use(sessions.Sessions("mysession", store), filter.GlobalFilter)
	engine.Use(sessions.Sessions("mysession", store), func(c *gin.Context) {
		s := sessions.Default(c)
		appSession := &session.AppSession{S: &s}
		filter.GlobalFilter(c, appSession, db)
	})
	// app.engine.GET("/prefright", func(c *gin.Context) {})
	controller.SetController(app, engine)
	return engine
}
