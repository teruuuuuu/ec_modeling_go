package filter

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"../session"
)

var GlobalFilter = func() func(c *gin.Context, as *session.AppSession, db *gorm.DB) {
	return func(c *gin.Context, as *session.AppSession, db *gorm.DB) {
		requestCount(c, as)
		loginCheck(c, as, db)
	}
}()
