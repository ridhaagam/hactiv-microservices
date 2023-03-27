package book

import domainBook "microservices/challenge-4-advance/domain/book"

func (book *Book) toDomainMapper() *domainBook.Book {
	return &domainBook.Book{
		ID:          book.ID,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		CreatedAt:   book.CreatedAt,
	}
}

func fromDomainMapper(book *domainBook.Book) *Book {
	return &Book{
		ID:          book.ID,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		CreatedAt:   book.CreatedAt,
	}
}

func arrayToDomainMapper(books *[]Book) *[]domainBook.Book {
	booksDomain := make([]domainBook.Book, len(*books))
	for i, book := range *books {
		booksDomain[i] = *book.toDomainMapper()
	}

	return &booksDomain
}
