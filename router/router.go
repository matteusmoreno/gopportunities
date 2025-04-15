package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := router.Run(":8090")
	if err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}
