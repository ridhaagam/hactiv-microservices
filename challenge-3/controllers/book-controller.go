package controllers

import (
	"microservices/challenge-3/repository/postgres/sqlc"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get Books
func (InPg *PgDB) GetBooks(ctx *gin.Context) {
	// Get books
	books, err := InPg.Master.GetBooks(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, books)
}

// Get Book
func (InPg *PgDB) GetBook(ctx *gin.Context) {
	// Check ID
	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Get book
	book, err := InPg.Master.GetBook(ctx, int64(bookID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, book)
}

// Add Book
func (InPg *PgDB) AddBook(ctx *gin.Context) {
	var (
		data sqlc.AddBookParams
		err  error
	)

	// Get json data
	err = ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book data"})
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

	// Add book
	book, err := InPg.Master.AddBook(ctx, data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

// Update Book
func (InPg *PgDB) UpdateBook(ctx *gin.Context) {
	var (
		data sqlc.UpdateBookParams
		err  error
	)

	// Check ID
	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Check book
	book, err := InPg.Master.GetBook(ctx, int64(bookID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Book not found"})
		return
	}

	// Get json data
	err = ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book data"})
		return
	}
	data.ID = book.ID

	// Update book
	err = InPg.Master.UpdateBook(ctx, data)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// Delete Book
func (InPg *PgDB) DeleteBook(ctx *gin.Context) {
	// Check ID
	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Check book
	_, err = InPg.Master.GetBook(ctx, int64(bookID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Book not found"})
		return
	}

	// Delete book
	err = InPg.Master.DeleteBook(ctx, int64(bookID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
