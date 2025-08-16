package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize router
	r := gin.Default()

	// Intialize Routes
	initializeRoutes(r)

	// Run server
	r.Run(":8081")
}
