package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	//server.Use(cors.Default())
	server.GET("/service1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.Run(":8081")
}
