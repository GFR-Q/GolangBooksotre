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
		&models.User{},
		&models.Book{},
		&models.Author{},
		&models.Category{},
		&models.FavoriteBook{},
	)

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/books", handlers.GetBooks)

	favRoutes := r.Group("/books")
	favRoutes.Use(handlers.AuthMiddleware())
	{
		favRoutes.GET("/favorites", handlers.GetFavorites)
		favRoutes.PUT("/:id/favorites", handlers.AddFavorite)
		favRoutes.DELETE("/:id/favorites", handlers.RemoveFavorite)
	}

	r.Run(":8080")
}
