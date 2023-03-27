package book

import (
	domainBook "microservices/challenge-4-advance/domain/book"
)

// NewBook is a struct that contains the data for a new book
type NewBook struct {
	Title       string `json:"title" example:"book title"`
	Author      string `json:"author" example:"mr. author"`
	Description string `json:"description" example:"book description"`
}

// PaginationResultBook is a struct that contains the pagination result for book
type PaginationResultBook struct {
	Data       *[]domainBook.Book
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
