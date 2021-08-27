package routes

import (
	"github.com/gin-gonic/gin"
	"golang-challenge/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine{
	main := router.Group("api/v1")
	{
		products := main.Group("products")
		{
			products.GET("/:id", controllers.ShowProduct)
		}
	}
	return router
}
