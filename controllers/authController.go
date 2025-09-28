package controllers

import (
	"context"
	"fmt"
	"net/http"
	"racer_http/models"
	"racer_http/repository"
	"racer_http/sqlite/entities"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	UserRepository repository.UserRepository
}

func NewAuthController(userRepository repository.UserRepository) *AuthController {
	return &AuthController{
		UserRepository: userRepository,
	}
}

func (controller *AuthController) CreateUser(c *gin.Context) {
	ctx := context.Background()

	var createUserRequest models.CreateUserRequest

	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userFound, err := controller.UserRepository.GetUserByEmail(ctx, createUserRequest.Email)
	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this email already exists"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createUserParams := entities.CreateUserParams{
		FirstName: createUserRequest.FirstName,
		LastName:  createUserRequest.LastName,
		Email:     createUserRequest.Email,
		Password:  string(passwordHash),
	}

	// TODO: handle error on create user
	user, _ := controller.UserRepository.CreateUser(ctx, createUserParams)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (controller *AuthController) Login(c *gin.Context) {
	ctx := context.Background()
	var authInput models.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userFound, _ := controller.UserRepository.GetUserByEmailForLogin(ctx, authInput.Email)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// TODO: Use a real secret from env variables for signing jwt
	token, err := generateToken.SignedString([]byte("SECRET"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.JSON(200, gin.H{"token": token})
}

func (controller *AuthController) CheckAuth(c *gin.Context) {
	ctx := context.Background()

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := authToken[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// TODO: Use a real secret from env variables for signing jwt
		return []byte("SECRET"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, _ := controller.UserRepository.GetUserById(ctx, int64(claims["id"].(float64)))
	if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set(
		"currentUser",
		user,
	)

	c.Next()
}

func (controller *AuthController) GetUser(c *gin.Context) {
	user, _ := c.Get("currentUser")
	c.JSON(200, gin.H{"user": user})
}
