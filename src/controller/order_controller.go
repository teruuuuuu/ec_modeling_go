package controller

import (
	"fmt"

	"../app"
	"../session"

	"github.com/gin-gonic/gin"
)

type ItemForm struct {
	ProductId uint `json:"product_id" binding:"required"`
	Number    uint `json:"number" binding:"required"`
}

func order_group(app *app.App) *ControllerGroup {
	var (
		order_controller = []Controller{}
	)
	addController := func(method int32, path string, handler func(c *gin.Context, as *session.AppSession)) {
		order_controller = append(order_controller, Controller{method: method, path: path, handler: handler})
	}
	addController(post, "/updateItem", func(c *gin.Context, as *session.AppSession) {
		itemForm := ItemForm{}
		c.BindJSON(&itemForm)
		_, loginUser := as.GetLoginUser()
		fmt.Println(app)
		fmt.Println(itemForm)
		ret := app.UpdateItem(loginUser.UserId, itemForm.ProductId, itemForm.Number)
		if ret {
			c.JSON(200, gin.H{
				"result": "success",
			})
		} else {
			c.JSON(200, gin.H{
				"result": "fail",
			})
		}
	})

	addController(get, "/cartItems", func(c *gin.Context, as *session.AppSession) {
		_, loginUser := as.GetLoginUser()
		c.JSON(200, gin.H{"result": app.CartItems(loginUser.UserId)})
	})

	return &ControllerGroup{group: "order", controllers: order_controller}
}
