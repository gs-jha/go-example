package rest

import "github.com/gin-gonic/gin"

func (a *App) SetRoutes() {
	router := gin.Default()

	router.GET("/long", LongRun)

	a.router = router
}
