package main

import (
	"github.com/gin-gonic/gin"
	racer "github.com/nickastewart/racer-parser"
)

func main() {
	router := gin.Default()
	router.GET("/parse", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": racer.Parse("/Users/nickstewart/results/2024-07-16-Milton-Keynes.eml"),
		})
	})

	router.Run()
}
