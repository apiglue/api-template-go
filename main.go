package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := SetupRouter()
	router.Run()
}

// SetupRouter -
func SetupRouter() *gin.Engine {

	router := gin.Default()

	v1 := router.Group("api/")
	{
		v1.GET("/heartbeat", GetHealthCheck)
	}

	return router
}

// GetHealthCheck - heartbeat endpoint
func GetHealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"ok": "healthy!"})
}
