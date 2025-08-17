package router

import (
	docs "github.com/DaviRodrigues/Project-go-with-gin/docs"
	"github.com/DaviRodrigues/Project-go-with-gin/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(r *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()

	const basePath string = "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := r.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.GET("/openings", handler.ListOpeningsHandler)

	}
	// Initialize Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
