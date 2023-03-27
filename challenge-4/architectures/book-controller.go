package architectures

import (
	"microservices/challenge-4/presentation"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get Books
func (InPg *PgDB) GetBooks(ctx *gin.Context) {
	// Get books
	books, err := GetBooks(InPg, 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, books)
}

// Get Book
func (InPg *PgDB) GetBook(ctx *gin.Context) {
	// Get book
	id := ctx.Param("id")
	book, err := GetBook(InPg, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Validate book
	if book.ID == 0 {
		ctx.JSON(http.StatusAccepted, gin.H{"message": "data tidak tersedia"})
		return
	}

	ctx.JSON(http.StatusAccepted, book)
}

// Add Book
func (InPg *PgDB) AddBook(ctx *gin.Context) {
	var (
		data presentation.Book
		err  error
	)

	// Get json data
	err = ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Validate book title
	if data.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title are required fields"})
		return
	}

	// Validate book author
	if data.Author == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Author are required fields"})
		return
	}

	// Add Book
	book, err := AddBook(InPg, data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

// Update Book
func (InPg *PgDB) UpdateBook(ctx *gin.Context) {
	var (
		data presentation.Book
		err  error
	)

	// Get book
	id := ctx.Param("id")
	book, err := GetBook(InPg, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Validate book
	if book.ID == 0 {
		ctx.JSON(http.StatusAccepted, gin.H{"message": "data tidak tersedia"})
		return
	}

	// Get json data
	err = ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Update Book
	err = UpdateBook(InPg, id, data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// Delete Book
func (InPg *PgDB) DeleteBook(ctx *gin.Context) {
	// Get book
	id := ctx.Param("id")
	book, err := GetBook(InPg, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Validate book
	if book.ID == 0 {
		ctx.JSON(http.StatusAccepted, gin.H{"message": "data tidak tersedia"})
		return
	}

	// Delete Book
	err = DeleteBook(InPg, book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
