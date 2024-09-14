package main

import (
	"books-backend/controllers"
	"books-backend/initialisers"
	"books-backend/middleware"

	"github.com/gin-gonic/gin"
)

// NEXT
// figure out whats up with the
// Make User Authentication & Authorisation (JWT)
// Research Deployment on AWS

// Better folder Structure --> MVC

func init() {
	initialisers.ConnectToDB()
	initialisers.SyncDatabase()
	initialisers.LoadEnvVariables()
}

// MAIN
func main() {
	router := gin.Default()

	router.SetTrustedProxies([]string{"109.81.95.132"})

	// Books
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBookByID)
	router.GET("/books/author/:author", controllers.GetBooksByAuthor)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.PUT("/books/read/:id", controllers.MarkRead)
	router.DELETE("/books/:id", controllers.DeleteBook)

	// Auth
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	router.Run(":8000")
}
