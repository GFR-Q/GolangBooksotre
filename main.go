package main

import (
	"bookstore/database"
	"bookstore/handlers"
	"bookstore/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(
		&models.Book{},
		&models.Author{},
		&models.Category{},
		&models.FavoriteBook{},
	)

	r := gin.Default()
	r.GET("/books", handlers.GetBooks)
	r.POST("/books", handlers.CreateBook)
	r.GET("/books/:id", handlers.GetBookByID)

	r.GET("/authors", handlers.GetAuthors)
	r.POST("/authors", handlers.CreateAuthor)
	favRoutes := r.Group("/books")
	favRoutes.Use(handlers.AuthMiddleware())
	{
		favRoutes.GET("/favorites", handlers.GetFavorites)
		favRoutes.PUT("/:id/favorites", handlers.AddFavorite)
		favRoutes.DELETE("/:id/favorites", handlers.RemoveFavorite)
	}

	r.Run(":8080")
}
