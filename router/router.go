package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	initializeRoutes(router)

	err := router.Run(":8090")
	if err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}
