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

	log.Info("zero say hello world 002")
	log.Info("Starting server on :8080")
	log.Info("leo say hi")

	log.Info("spring is sp_002_leo 003")
	log.Info("spring is sp_002_dylan.su 002") // dylan.su@bollertech.com
	r.Run(":8080")                            // 監聽 :8080
}
