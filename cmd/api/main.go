package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// course video link https://www.youtube.com/watch?v=S069igHKUIw&t=2233s

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"answer": "pong",
		})
	})

	router.Run()
}
