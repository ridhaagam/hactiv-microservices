package main

import (
	"microservices/challenge-4/architectures"
	"microservices/challenge-4/repository/postgres"
	"microservices/challenge-4/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Postgres
	pg := postgres.New()
	InPg := &architectures.PgDB{Master: pg}

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
