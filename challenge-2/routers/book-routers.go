package routers

import (
	"github.com/gin-gonic/gin"
	"microservices/challenge-2/controllers"
)

func BookRouters(router *gin.Engine) {
	RBook := router.Group("/books")
	{
		RBook.GET("", controllers.GetBooks)
		RBook.GET("/:id", controllers.GetBook)
		RBook.POST("", controllers.AddBook)
		RBook.PUT("/:id", controllers.UpdateBook)
		RBook.DELETE("/:id", controllers.DeleteBook)
	}
}
