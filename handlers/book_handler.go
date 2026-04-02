package handlers

import (
	"bookstore/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{
	{ID: 1, Title: "The Go Programming Language", AuthorID: 1, CategoryID: 1, Price: 35.50},
	{ID: 2, Title: "Clean Code", AuthorID: 2, CategoryID: 2, Price: 42.00},
}
var nextBookID = 3

func GetBooks(c *gin.Context) {
	categoryIDStr := c.Query("category_id")
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	filteredBooks := books
	if categoryIDStr != "" {
		categoryID, _ := strconv.Atoi(categoryIDStr)
		var temp []models.Book
		for _, b := range books {
			if b.CategoryID == categoryID {
				temp = append(temp, b)
			}
		}
		filteredBooks = temp
	}

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	if startIndex > len(filteredBooks) {
		c.JSON(http.StatusOK, []models.Book{})
		return
	}
	if endIndex > len(filteredBooks) {
		endIndex = len(filteredBooks)
	}

	c.JSON(http.StatusOK, filteredBooks[startIndex:endIndex])
}

func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if newBook.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be positive"})
		return
	}

	newBook.ID = nextBookID
	nextBookID++
	books = append(books, newBook)

	c.JSON(http.StatusCreated, newBook)
}

func GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, b := range books {
		if b.ID == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}
