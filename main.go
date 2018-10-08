package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := SetupRouter()
	router.Run(":" + port)
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
