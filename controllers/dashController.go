package controllers

import (
	"books-backend/initialisers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PercentageRead(c *gin.Context) {
	db := initialisers.ConnectToDB()

	var body struct {
		User_id int
	}

	c.BindJSON(&body)

	var total_count int

	// fetch total books of user
	result := db.Raw(`
    SELECT COUNT(*)
	FROM user_books
	WHERE user_id = 2;
	`, body.User_id).
		Find(&total_count) // Use Scan for multiple records
	if result.Error != nil {
		c.JSON(http.StatusNoContent, gin.H{"message": "Database issue."})
	}

	fmt.Println(total_count)

	// fetch read books by user
	var read int

	result = db.Raw(`
    SELECT COUNT(*)
	FROM user_books
	WHERE user_id = 2 AND status = 1;
	`, body.User_id).
		Find(&read) // Use Scan for multiple records
	if result.Error != nil {
		c.JSON(http.StatusNoContent, gin.H{"message": "Database issue."})
	}

	fmt.Println(read)

	percentage := (float64(read) / float64(total_count)) * 100

	// return percentage
	fmt.Println("Getting Percentage Read.")
	c.JSON(http.StatusOK, gin.H{"Percentage Read": percentage})

}
