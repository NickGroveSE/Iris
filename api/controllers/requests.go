package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetServerHealth(c *gin.Context) {
	var Response struct {
		Health string
	}
	Response.Health = "I am healthy!"
	c.JSON(http.StatusOK, Response)
}

func GetMusicData(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"response": "hello"})

}
