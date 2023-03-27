package main

import (
	"microservices/challenge-3/controllers"
	"microservices/challenge-3/repository/postgres"
	"microservices/challenge-3/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Postgres
	pg := postgres.New()
	InPg := &controllers.PgDB{Master: pg}

	// Init Router
	router := gin.Default()

	// Books Router
	routers.BookRouters(router, InPg)

	// Router Port
	err := router.Run(":4000")
	if err != nil {
		panic("Error When Running")
	}
}
