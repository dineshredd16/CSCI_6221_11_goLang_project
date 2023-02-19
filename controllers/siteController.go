package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SitesCreate (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}