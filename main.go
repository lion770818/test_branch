package main

import (
	"github.com/Allenxuxu/gev/log"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// GET /ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Info("zero say hello world")
	log.Info("Starting server on :8080")

	r.Run(":8080") // 監聽 :8080
}
