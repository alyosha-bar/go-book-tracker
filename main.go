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
	initialisers.SyncDatabase()
	initialisers.ConnectToDB()
	initialisers.LoadEnvVariables()
}

// MAIN
func main() {
	initialisers.SyncDatabase()
	router := gin.Default()

	router.SetTrustedProxies([]string{"109.81.95.132"})

	// Books
	// seems to be okay
	router.GET("/books", controllers.GetBooks) //✅

	// both need slight changes
	router.GET("/books/:id", controllers.GetBookByID)
	router.GET("/books/author/:author", controllers.GetBooksByAuthor)
	// seems to be okay
	router.POST("/books", controllers.CreateBook) // ✅

	// Needs alteration
	router.PUT("/books/:id", controllers.UpdateBook)

	// Needs a lot of work
	router.PUT("/books/read", controllers.MarkRead)

	// Easy Work
	router.DELETE("/books/:id", controllers.DeleteBook)

	// Auth
	router.POST("/signup", controllers.SignUp)                            // ✅
	router.POST("/login", controllers.Login)                              // ✅
	router.GET("/validate", middleware.RequireAuth, controllers.Validate) // ✅

	// Dashboard Routes:
	// Fetch Percentage Read
	router.GET("/dash/percentage", controllers.PercentageRead)
	// Fetch read of the past week, month and year
	// Count most common authors

	router.Run(":8000")
}
