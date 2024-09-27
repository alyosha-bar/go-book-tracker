package controllers

import (
	"books-backend/initialisers"
	"books-backend/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	var body struct {
		User_id  int
		Username string
		Email    string
		Password string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": " Bad request"})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to Hash"})

		return
	}

	user := models.User{User_id: body.User_id, Username: body.Username, Email: body.Email, Password: string(hash)}
	result := initialisers.ConnectToDB().Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error adding user"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Signed Up."})

	fmt.Println("Signing up")
}

func Login(c *gin.Context) {

	fmt.Println("Logging in on the server ... ")

	var body struct {
		Email    string
		Password string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": " Bad request."})
		return
	}

	// validation logic
	// fetch user where email == body.Email
	DB := initialisers.ConnectToDB()
	var result models.User

	// Use DB query to fetch the user, no need to assign the query result to 'result'
	err := DB.Model(&models.User{}).Where("Email = ?", body.Email).First(&result).Error

	// Check if there is an error (e.g., user not found)
	if err != nil {
		// Handle error (e.g., user not found, DB connection issue)
		c.JSON(http.StatusBadRequest, gin.H{"message": "User does not exist."})
		fmt.Println("Error: ", err)
		return
	}

	// Compare password and password hash
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password Incorrect."})
		return
	}

	//respond with a JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		"sub": result.User_id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET"))) // replace with ENV variable

	// fmt.Println(os.Getenv("SECRET"))

	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to Create Token."})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true) // look into these settings

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged In. Responding with JWT.",
	})
	fmt.Println("Loggin In")
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	// gives access to user object from the request
	c.JSON(http.StatusOK, gin.H{"message": user})
}
