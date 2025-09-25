package main

import (
	"database/sql"
	_ "embed"
	"github.com/gin-gonic/gin"
	racer "github.com/nickastewart/racer-parser"
	"log"
	_ "modernc.org/sqlite"
	"racer_http/sqlite/entities"
	"racer_http/repository"
	"context"
)

func main() {

	ctx := context.Background()

	db, err := sql.Open("sqlite", "/Users/nickstewart/sqlite/racer.db")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	queries := entities.New(db)
	var userRepository repository.UserRepository = repository.NewUserRepository(queries)

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

	router.GET("/user", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"User": userRepository.GetUserById(ctx, int64(1)),
		})
	})

	router.Run()
}
