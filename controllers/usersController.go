package controllers

import (
	"books-backend/initialisers"
	"books-backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	var body struct {
		Username string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": " Bad request"})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Hash"})

		return
	}

	user := models.User{Username: body.Username, Email: body.Email, Password: string(hash)}
	result := initialisers.ConnectToDB().Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error adding user"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Signed Up."})

	fmt.Println("Signing up")
}

func Login(c *gin.Context) {
	fmt.Println("Loggin In")
}
