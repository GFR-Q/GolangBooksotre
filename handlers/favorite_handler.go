package handlers

import (
	"bookstore/database"
	"bookstore/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddFavorite(c *gin.Context) {
	userID, _ := c.Get("user_id")
	bookID, _ := strconv.Atoi(c.Param("id"))

	fav := models.FavoriteBook{
		UserID: uint(userID.(float64)),
		BookID: uint(bookID),
	}

	if err := database.DB.Create(&fav).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Уже в избранном"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Added to favorites"})
}
func RemoveFavorite(c *gin.Context) {
	userID, _ := c.Get("user_id")
	bookID, _ := strconv.Atoi(c.Param("id"))

	database.DB.Where("user_id = ? AND book_id = ?", userID, bookID).Delete(&models.FavoriteBook{})
	c.JSON(http.StatusOK, gin.H{"message": "Removed"})
}
func GetFavorites(c *gin.Context) {
	userID, _ := c.Get("user_id")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var books []models.Book
	database.DB.Table("books").
		Joins("JOIN favorite_books ON favorite_books.book_id = books.id").
		Where("favorite_books.user_id = ?", userID).
		Limit(limit).Offset(offset).
		Find(&books)

	c.JSON(http.StatusOK, books)
}
