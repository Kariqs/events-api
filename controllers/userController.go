package controllers

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/Kariqs/events-api/initializers"
	"github.com/Kariqs/events-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignUp(ctx *gin.Context) {
	//get user data from request body.
	var userData models.User
	ctx.ShouldBindJSON(&userData)

	//Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)
	if err != nil {
		panic(err)
	}
	userData.Password = string(hashedPassword)
	//create user
	result := initializers.DB.Create(&userData)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "user with this email already exists."})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "error creating user"})
		return
	}

	//return created user
	response := map[string]any{"message": "user created successfully"}
	ctx.JSON(http.StatusCreated, response)
}

func Login(ctx *gin.Context) {
	//Get email and password from the body
	var loginData models.LoginData
	ctx.ShouldBindJSON(&loginData)

	//Look up for the user to see if they exist
	var user models.User
	result := initializers.DB.First(&user, "email=?", loginData.Email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusCreated, gin.H{"error": "invalid email or password"})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"error": "unable to look up for user"})
		return
	}

	//Compare the password they sent
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		ctx.JSON(http.StatusCreated, gin.H{"error": "invalid email or password"})
		return
	}

	//Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	jwtSecret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unable to login"})
		panic(err)
	}

	//Send the Token
	response := map[string]any{
		"token":   tokenString,
		"message": "Login was successful",
	}

	ctx.JSON(http.StatusOK, response)
}
