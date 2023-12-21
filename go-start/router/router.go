package router

import (
	"go/api/go-start/controller"

	"github.com/gin-gonic/gin"
)

func DefineRoute(router *gin.Engine) {
	products := router.Group("/products")
	{
		products.GET("/", controller.GetAllProducts)
		products.GET("/:id", controller.GetProductById)
		products.POST("/", controller.PostProduct)
	}
}
