package controllers

import (
	"net/http"
	"os"
	"time"
	"webScraperBackend/initializers"
	"webScraperBackend/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context){

	var body struct {
		Email string
		Password string
	}

	if (c.Bind(&body) != nil) {
		c.JSON(http.StatusBadRequest, gin.H {
			"status": false,
			"message": "failed to read the body",
			"data": "",
		})
		return 
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"status": false,
			"message": "failed to hash the password, " + err.Error(),
			"data": "",
		})
		return 
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"status": false,
			"message": "unable to create the user, " + result.Error.Error(),
			"data": "",
		})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"messsage": "user created succesfully",
		"data": user,
	})
}

func Login(c *gin.Context){
	var body struct {
		Email string
		Password string
	}
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H {
			"status": false,
			"message": "invalid credentials",
			"data": "",
		})
		return 
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if (err != nil) {
		c.JSON(http.StatusBadRequest, gin.H {
			"status": false,
			"message": "password invalid credentials",
			"data": "",
		})
		return 
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.ID,
		"expiry": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY_DEV")))
	if (err != nil) {
		c.JSON(http.StatusBadRequest, gin.H {
			"status": false,
			"message": "invalid credentials",
			"data": "",
		})
		return 
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"messsage": "login succesful",
		"data": tokenString,
	})
}