package main

import (
	"net/http"
	"webScraperBackend/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
  initializers.LoadEnvVariables()
  initializers.ConnectToDatabase()
}

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "Hello World!",
    })
  })
  r.Run()
}