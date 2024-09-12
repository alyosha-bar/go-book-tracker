package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NEXT
// Make User Authentication & Authorisation (JWT)
// Research Deployment on AWS

type Book struct {
	id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Read   int    `json:read`
}

type Req struct {
	read int `json:"value"`
}

// GET functions

// all books --> FROM DATABASE!!
func getBooks(c *gin.Context) {

	// SINCE I HAVE READ AND UNREAD I CAN USE GO ROUTINES TO FETCH BOTH AND RETURN THEM SEPARATELY TO HELP WITH FRONTEND

	db := connectToDB()

	var books []Book
	fmt.Println("Fetching first book...")
	var result = db.Find(&books)
	if result.Error != nil {
		panic("Failed to fetch data: " + result.Error.Error())
	}
	fmt.Println("Fetched book:", books)

	c.JSON(http.StatusOK, books)
}

// get books by ID --> FROM DATABASE!!
func getBookByID(c *gin.Context) {

	db := connectToDB()

	id := c.Param("id")

	var book Book
	var result = db.First(&book, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "book not found :("})
		panic("Failed to fetch book: " + result.Error.Error())

	}

	c.JSON(http.StatusOK, book)
}

// get books by author
func getBooksByAuthor(c *gin.Context) {

	db := connectToDB()

	id := c.Param("author")

	var books []Book
	var result = db.Find(&books, "author = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "author has no books :("})
		panic("Failed to fetch book: " + result.Error.Error())

	}

	c.JSON(http.StatusOK, books)
}

// POST --> INSERT STATEMENT --> TO DATABASE!!
func createBook(c *gin.Context) {

	db := connectToDB()

	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Inserting data...")
	result := db.Create(newBook)
	if result.Error != nil {
		panic("Failed to insert data: " + result.Error.Error())
	}
	fmt.Println("Data inserted")

	// books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

// PUT --> UPDATE STATEMENT --> TO DATABASE!!
func updateBook(c *gin.Context) {

	db := connectToDB()

	var updateBook Book

	if err := c.BindJSON(&updateBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Inserting data...")

	// might need some figuring out how to optimise
	result := db.Save(updateBook)
	if result.Error != nil {
		panic("Failed to insert data: " + result.Error.Error())
	}
	fmt.Println("Data inserted")

	// books = append(books, newBook)
	c.JSON(http.StatusCreated, updateBook)

}

func markRead(c *gin.Context) {

	db := connectToDB()
	id := c.Param("id")

	var body Req

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Inserting data...")

	fmt.Printf("Value: %v", body.read)
	fmt.Println("")

	var updateBook Book
	// might need some figuring out how to optimise
	result := db.Model(&updateBook).Where("id = ?", id).UpdateColumn("Read", gorm.Expr("?", body.read))
	if result.Error != nil {
		panic("Failed to insert data: " + result.Error.Error())
	}
	fmt.Println("Data inserted")

	// books = append(books, newBook)
	c.JSON(http.StatusCreated, gin.H{"message": "Marked Read"})

}

// DELETE --> DELETE
func deleteBook(c *gin.Context) {
	id := c.Param("id")

	fmt.Print(id)

	c.JSON(http.StatusNotFound, gin.H{"message": "book not found :("})

}

// UTIL --> FUNCTIONS
func connectToDB() *gorm.DB {
	fmt.Println("Connecting to database...")

	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to Database: " + err.Error())
	}
	fmt.Println("Connected to database")

	// fmt.Println("Migrating schema...")
	// err = db.AutoMigrate(&Book{})
	// if err != nil {
	// 	panic("Failed to migrate schema: " + err.Error())
	// }
	// fmt.Println("Schema migrated")

	return db
}

// MAIN
func main() {
	router := gin.Default()

	router.SetTrustedProxies([]string{"109.81.95.132"})

	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.GET("/books/author/:author", getBooksByAuthor)
	router.POST("/books", createBook)
	router.PUT("/books/:id", updateBook)
	router.PUT("/books/read/:id", markRead)
	router.DELETE("/books/:id", deleteBook)

	router.Run(":8000")
}
