package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMusicData(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"response": "hello"})

}
