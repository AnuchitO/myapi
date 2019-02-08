package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	log.Println("in hello handler")
	c.JSON(http.StatusOK, gin.H{"status": "hello"})
}

func loginMiddleware(c *gin.Context) {
	log.Println("starting middleware")
	authKey := c.GetHeader("Authorization")
	if authKey != "token123" {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}
	c.Next()

	log.Println("ending middleware")
}

func main() {
	r := gin.Default()

	r.Use(loginMiddleware)

	r.POST("/hello", helloHandler)
	r.Run(":1234")
}
