package main

import (
	"github.com/gin-gonic/gin"
	racer "github.com/nickastewart/racer-parser"
)

func main() {
	router := gin.Default()
	router.GET("/parse", func(c *gin.Context) {
		filePath := c.Query("file")
		c.JSON(200, gin.H{
			"message": racer.Parse(filePath),
		})
	})

	router.Run()
}
