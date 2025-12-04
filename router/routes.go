package router

import (
	"github.com/felipematheus1337/Gopportunities/handler"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/override/docs"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{

		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.GET("/opening", handler.ListOpeningHandler)

	}

}
