package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	router := SetupRouter()
	router.Run()
}

// SetupRouter -
func SetupRouter() *gin.Engine {

	router := gin.Default()

	v1 := router.Group("/")
	{
		v1.GET("/health", GetHealthCheck)
	}

	return router
}

// GetHealthCheck - heartbeat endpoint
func GetHealthCheck(c *gin.Context) {

	content := gin.H{
		"id":          strconv.Itoa(os.Getpid()),
		"description": "PID says hello!",
	}

	c.JSON(200, content)
}
