package rest

import "github.com/gin-gonic/gin"

func SetRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/", HelloWorld)

	return router
}
