package filter

import (
	"fmt"

	"../session"
	"github.com/gin-gonic/gin"
)

func requestCount(c *gin.Context, as *session.AppSession) {
	isValid, count := as.GetCount()
	if !isValid {
		count = 0
	}
	count++
	fmt.Println("count: ", count)
	as.SetCount(count)
}
