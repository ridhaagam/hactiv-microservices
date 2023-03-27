// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	bookService "microservices/challenge-4-advance/application/usecases/book"
	bookRepository "microservices/challenge-4-advance/infrastructure/repository/postgres/book"
	bookController "microservices/challenge-4-advance/infrastructure/rest/controllers/book"

	"gorm.io/gorm"
)

// BookAdapter is a function that returns a book controller
func BookAdapter(db *gorm.DB) *bookController.Controller {
	mRepository := bookRepository.Repository{DB: db}
	service := bookService.Service{BookRepository: mRepository}
	return &bookController.Controller{BookService: service}
}
