package controllers

import (
	"net/http"
	"strconv"

	svc "github.com/VikashChauhan51/go-sample-api/internal/core/interfaces/services"
	"github.com/VikashChauhan51/go-sample-api/internal/dto"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService svc.BookService
}

func NewBookController(service svc.BookService) *BookController {
	return &BookController{service}
}

// @Summary Get books
// @Description Retrieves a list of books
// @Tags books
// @Produce  json
// @Success 200 {array} dto.Book
// @Router /books [get]
func (b *BookController) GetBooks(c *gin.Context) {
	books, err := b.bookService.FetchBooksAsync()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// @Summary Create book
// @Description Creates a new book
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body dto.BookCreateDTO true "Book to create"
// @Success 201 {object} dto.Book
// @Router /books [post]
func (b *BookController) CreateBook(c *gin.Context) {
	var bookDTO dto.BookCreateDTO
	if err := c.ShouldBindJSON(&bookDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := b.bookService.CreateBook(&bookDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully"})
}

// @Summary Get book by ID
// @Description Retrieves a book by its ID
// @Tags books
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} dto.Book
// @Router /book/{id} [get]
func (b *BookController) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := b.bookService.FetchBookByID(bookID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// @Summary Update book
// @Description Updates an existing book
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Param book body dto.BookUpdateDTO true "Book to update"
// @Success 200 {object} dto.Book
// @Router /books/{id} [put]
func (b *BookController) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var bookDTO dto.BookUpdateDTO
	if err := c.ShouldBindJSON(&bookDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bookDTO.ID = bookID

	err = b.bookService.UpdateBook(&bookDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

// @Summary Delete book
// @Description Deletes a book by its ID
// @Tags books
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {string} string "Book deleted successfully"
// @Router /book/{id} [delete]
func (b *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	err = b.bookService.DeleteBook(bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
