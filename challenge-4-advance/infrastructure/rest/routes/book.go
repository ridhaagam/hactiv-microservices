// Package routes contains all routes of the application
package routes

import (
	bookController "microservices/challenge-4-advance/infrastructure/rest/controllers/book"

	"github.com/gin-gonic/gin"
)

// UserRoutes is a function that contains all routes of the book
func BookRoutes(router *gin.RouterGroup, controller *bookController.Controller) {
	routerBook := router.Group("/books")
	// routerBook.Use(middlewares.AuthJWTMiddleware())
	{
		routerBook.GET("", controller.GetAllBooks)
		routerBook.GET("/:id", controller.GetBookByID)
		routerBook.POST("", controller.NewBook)
		routerBook.PUT("/:id", controller.UpdateBook)
		routerBook.DELETE("/:id", controller.DeleteBook)
	}
}
