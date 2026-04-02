package handlers

import (
	"bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var categories = []models.Category{
	{ID: 1, Name: "Action"},
	{ID: 2, Name: "Horror"},
}
var nextCategoryID = 3

func GetCategory(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	category.ID = nextCategoryID
	nextCategoryID++
	categories = append(categories, category)
	c.JSON(http.StatusCreated, category)
}
