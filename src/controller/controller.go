package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"../app"
	"../session"
)

const (
	get = iota
	post
	put
)

type Controller struct {
	method  int32
	path    string
	handler func(c *gin.Context, s *session.AppSession)
}

type ControllerGroup struct {
	group       string
	controllers []Controller
}

func SetController(app *app.App, engine *gin.Engine) {
	product_group(app).addController(engine)
	order_group(app).addController(engine)
	// h_group.addController(engine)
	// d_group.addController(engine)
}

func (controllerGroup *ControllerGroup) addController(engine *gin.Engine) {
	group := engine.Group(controllerGroup.group)
	// group.Use(controllerGroup.filter)
	for i := 0; i < len(controllerGroup.controllers); i++ {
		controllerGroup.controllers[i].add(group)
	}
}

func (controller *Controller) add(group *gin.RouterGroup) {
	var method func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	if controller.method == get {
		method = group.GET
	} else if controller.method == post {
		method = group.POST
	} else if controller.method == put {
		method = group.PUT
	}

	method(controller.path, func(c *gin.Context) {
		s := sessions.Default(c)
		appSession := &session.AppSession{S: &s}
		controller.handler(c, appSession)
	})
}
