package architectures

import (
	"microservices/challenge-4/presentation"
)

// Get Books
func GetBooks(InPg *PgDB, limit int) (data []presentation.Book, err error) {
	respond := InPg.Master.Table("books").Limit(limit).Find(&data)
	if respond.Error != nil {
		return data, respond.Error
	}

	return data, nil
}

// Get Book
func GetBook(InPg *PgDB, id string) (data presentation.Book, err error) {
	respond := InPg.Master.Table("books").Where("id = ?", id).Find(&data)
	if respond.Error != nil {
		return data, respond.Error
	}

	return data, nil
}

// Add Book
func AddBook(InPg *PgDB, data presentation.Book) (presentation.Book, error) {
	respond := InPg.Master.Table("books").Create(&data)
	if respond.Error != nil {
		return data, respond.Error
	}

	return data, nil
}

// Update Book
func UpdateBook(InPg *PgDB, id string, data presentation.Book) error {
	respond := InPg.Master.Table("books").Where("id = ?", id).Updates(&data)
	if respond.Error != nil {
		return respond.Error
	}

	return nil
}

// Delete Book
func DeleteBook(InPg *PgDB, data presentation.Book) error {
	respond := InPg.Master.Table("books").Delete(&data)
	if respond.Error != nil {
		return respond.Error
	}

	return nil
}
