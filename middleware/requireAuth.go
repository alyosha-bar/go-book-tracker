package middleware

import (
	"books-backend/initialisers"
	"books-backend/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("In Middleware.")

	// Get cookie off req
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// decode and validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: &v", token.Header["alg"])
		}

		// return []byte(os.Getenv("SECRET")), nil
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check the exp
		fmt.Println("Made it into the func.")

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		fmt.Println("HERE NOW.")

		// find the user w the token sub
		DB := initialisers.ConnectToDB()
		var user models.User
		DB.First(&user, claims["sub"])

		if user.User_id == 0 {
			fmt.Println("User id is 0.")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// attach to req
		c.Set("user", user)

		//continue
		c.Next()
	} else {
		fmt.Println("We are here.")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}
