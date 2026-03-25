package handlers

import (
	"bookstore/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var authors = []models.Author{
	{ID: 1, Name: "Alan Donovan"},
	{ID: 2, Name: "Robert Martin"},
}
var nextAuthorID = 3

func GetAuthors(c *gin.Context) {
	c.JSON(http.StatusOK, authors)
}

func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	author.ID = nextAuthorID
	nextAuthorID++
	authors = append(authors, author)
	c.JSON(http.StatusCreated, author)
}
