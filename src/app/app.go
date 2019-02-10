package app

import (
	"github.com/jinzhu/gorm"
)

type App struct {
	db *gorm.DB
}

func New(db *gorm.DB) *App {
	app := &App{db: db}
	app.init()
	return app
}

func (app *App) init() {
	initData(app.db)
}
