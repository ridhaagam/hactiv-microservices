package book

import (
	domainBook "microservices/challenge-4-advance/domain/book"
	"time"

	"gorm.io/gorm"
)

// Book is a struct that contains the book model
type Book struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" example:"book title"`
	Author      string    `json:"author" example:"mr. author"`
	Description string    `json:"description" example:"book description"`
	CreatedAt   time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt   *gorm.DeletedAt
}

// TableName overrides the table name used by User to `users`
func (*Book) TableName() string {
	return "books"
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
