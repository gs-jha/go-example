package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LongRun(c *gin.Context) {
	fmt.Println("Starting long request")
	go longProcess(c)
	select {
	case <-c.Request.Context().Done():
		fmt.Println("Process canceled")
	}
	c.JSON(http.StatusOK, gin.H{"data": "Long process"})
}

func longProcess(c *gin.Context) {
	time.Sleep(20 * time.Second)

	fmt.Println("Task complete")
}
