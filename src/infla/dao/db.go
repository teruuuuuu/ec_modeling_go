package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db, err = func() (*gorm.DB, error) {
	var db, err = gorm.Open("sqlite3", "./tmp/gorm.db")
	if err != nil {
		fmt.Errorf("invalid database source: %v is not a valid type", err)
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	initDb(db)
	return db, err
}()

func GetDb() *gorm.DB {
	return db
}
