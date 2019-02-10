package controller

import (
	"../app"
	"../session"

	"github.com/gin-gonic/gin"
)

func product_group(app *app.App) *ControllerGroup {
	var (
		product_controller = []Controller{}
	)
	addController := func(method int32, path string, handler func(c *gin.Context, as *session.AppSession)) {
		product_controller = append(product_controller, Controller{method: method, path: path, handler: handler})
	}
	addController(get, "/search", func(c *gin.Context, as *session.AppSession) {
		c.JSON(200, gin.H{
			"result": app.ShowProducts(),
		})
	})

	return &ControllerGroup{group: "product", controllers: product_controller}
}
