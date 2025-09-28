package main

import (
	"database/sql"
	_ "embed"
	"github.com/gin-gonic/gin"
	racer "github.com/nickastewart/racer-parser"
	"log"
	_ "modernc.org/sqlite"
	"racer_http/controllers"
	"racer_http/repository"
	"racer_http/sqlite/entities"
)

func main() {

	db, err := sql.Open("sqlite", "/Users/nickstewart/sqlite/racer.db")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	queries := entities.New(db)
	var userRepository repository.UserRepository = repository.NewUserRepository(queries)
	authController := controllers.NewAuthController(userRepository)
	if err != nil {
		log.Panic(err)
	}
	router := gin.Default()

	router.GET("/parse", func(c *gin.Context) {
		filePath := c.Query("file")
		c.JSON(200, gin.H{
			"message": racer.Parse(filePath),
		})
	})

	router.POST("auth/user", authController.CreateUser)
	router.GET("auth/user", authController.Login)
	router.GET("user", authController.CheckAuth, authController.GetUser)

	router.Run()
}
