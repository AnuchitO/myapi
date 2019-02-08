package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func main() {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	r.Run() // listen and serve on 127.0.0.0:8080
}
