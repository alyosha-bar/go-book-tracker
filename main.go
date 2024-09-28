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
	router.GET("/api/books/:userid", controllers.GetBooks) //✅

	// both need slight changes
	router.GET("/api/books/id/:id", controllers.GetBookByID) // ✅
	router.GET("/api/books/author/:author", controllers.GetBooksByAuthor)
	// seems to be okay
	router.POST("/api/books", controllers.CreateBook) // ✅

	// Needs alteration
	router.PUT("/api/books/update/:id", controllers.UpdateBook)
	router.PUT("/api/books/read", controllers.MarkRead) // ✅

	// Easy Work
	router.DELETE("/api/books/:id", controllers.DeleteBook) // ✅

	// Auth
	router.POST("/api/signup", controllers.SignUp)                            // ✅
	router.POST("/api/login", controllers.Login)                              // ✅
	router.GET("/api/validate", middleware.RequireAuth, controllers.Validate) // ✅

	// Dashboard Routes:
	// Fetch Percentage Read
	router.GET("/api/dash/percentage", controllers.PercentageRead) // ✅
	// Fetch read of the past week, month and year
	// Count most common authors

	router.Run(":8000")
}
