package book

// Package medicine provides the use case for medicine

import (
	domainBook "microservices/challenge-4-advance/domain/book"
)

func (n *NewBook) toDomainMapper() *domainBook.Book {
	return &domainBook.Book{
		Title:       n.Title,
		Author:      n.Author,
		Description: n.Description,
	}
}
