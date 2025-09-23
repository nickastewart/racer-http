package main

import (
	"context"
	"database/sql"
	_ "embed"
	"github.com/gin-gonic/gin"
	racer "github.com/nickastewart/racer-parser"
	"log"
	_ "modernc.org/sqlite"
	"racer_http/racer-http-db"
)

func main() {

	ctx := context.Background()

	db, err := sql.Open("sqlite", "/Users/nickstewart/sqlite/racer.db")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	queries := racer_http.New(db)

	user, err := queries.GetUser(ctx, int64(1))

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
			"User": user,
		})
	})

	router.Run()
}

