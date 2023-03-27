package routers

import (
	"github.com/gin-gonic/gin"
	"microservices/challenge-3/controllers"
)

func BookRouters(router *gin.Engine, InPg *controllers.PgDB) {
	RBook := router.Group("/books")
	{
		RBook.GET("", InPg.GetBooks)
		RBook.GET("/:id", InPg.GetBook)
		RBook.POST("", InPg.AddBook)
		RBook.PUT("/:id", InPg.UpdateBook)
		RBook.DELETE("/:id", InPg.DeleteBook)
	}
}
