package rest

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func NewApp() App {
	app := App{}
	app.SetRoutes()

	return app
}

func (app App) Run(port string) {
	app.router.Run(port)
}
