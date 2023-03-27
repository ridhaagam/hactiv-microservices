package routers

import (
	"microservices/challenge-4/architectures"

	"github.com/gin-gonic/gin"
)

func BookRouters(router *gin.Engine, InPg *architectures.PgDB) {
	routerBook := router.Group("/books")
	{
		routerBook.GET("", InPg.GetBooks)
		routerBook.GET("/:id", InPg.GetBook)
		routerBook.POST("", InPg.AddBook)
		routerBook.PUT("/:id", InPg.UpdateBook)
		routerBook.DELETE("/:id", InPg.DeleteBook)
	}
}
