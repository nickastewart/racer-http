package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/parse", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "parse",
		})
	})

	router.Run()
}
