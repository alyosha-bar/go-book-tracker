package controllers

import (
	"books-backend/initialisers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Book struct {
	Book_id int
	Title   string `json:"title"`
	Author  string `json:"author"`
	Read    int    `json:read`
}

type Req struct {
	read int `json:"value"`
}

// GET functions

// all books BY USER
func GetBooks(c *gin.Context) {
	db := initialisers.ConnectToDB()

	var body struct {
		User_id int
	}

	c.BindJSON(&body)

	// save into this array
	type BookUserStatus struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		Status int    `json:"status"`
	}

	var books []BookUserStatus

	// db = db.Debug()

	result := db.Raw(`
    SELECT books.title, books.author, user_books.status
    FROM books
    INNER JOIN user_books ON user_books.book_id = books.book_id
    WHERE user_books.user_id = ?`, body.User_id).
		Find(&books) // Use Scan for multiple records

	if result.Error != nil {
		panic("Failed to fetch data: " + result.Error.Error())
	}

	fmt.Println(books)

	c.JSON(http.StatusOK, books)
}

// get books by ID --> FROM DATABASE!! --> AND BY ID
func GetBookByID(c *gin.Context) {

	db := initialisers.ConnectToDB()

	id := c.Param("id")

	var book Book
	var result = db.First(&book, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "book not found :("})
		panic("Failed to fetch book: " + result.Error.Error())

	}

	c.JSON(http.StatusOK, book)
}

// get books by author --> AND BY ID
func GetBooksByAuthor(c *gin.Context) {

	db := initialisers.ConnectToDB()

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
func CreateBook(c *gin.Context) {

	db := initialisers.ConnectToDB()

	var body struct {
		User_id int
		Book_id int
		Title   string
		Author  string
	}

	type UserBook struct {
		User_id int
		Book_id int
		Status  int
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println()

	var newBook = Book{Book_id: body.Book_id, Title: body.Title, Author: body.Author}

	var newUserBook = UserBook{User_id: body.User_id, Book_id: body.Book_id, Status: 0}

	fmt.Println(newBook)
	fmt.Println(newUserBook)

	fmt.Println("Inserting Book...")
	result := db.Create(newBook)
	if result.Error != nil {
		panic("Failed to insert data: " + result.Error.Error())
	}
	fmt.Println("Data inserted")

	fmt.Println("Inserting User Book Relation...")
	result = db.Create(newUserBook)
	if result.Error != nil {
		panic("Failed to insert data: " + result.Error.Error())
	}
	fmt.Println("Data inserted")

	// books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

// PUT --> UPDATE STATEMENT --> TO DATABASE!!
// REVIEW ENTIRE FUNCTION --> non priority
func UpdateBook(c *gin.Context) {

	db := initialisers.ConnectToDB()

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

// needs changing with new STATUS system
func MarkRead(c *gin.Context) {

	db := initialisers.ConnectToDB()
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
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	fmt.Print(id)

	c.JSON(http.StatusNotFound, gin.H{"message": "book not found :("})

}
