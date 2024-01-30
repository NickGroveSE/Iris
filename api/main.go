package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	app.GET("/health", func(c *gin.Context) {
		var Response struct {
			Health string
		}
		Response.Health = "I am healthy!"
		c.JSON(http.StatusOK, Response)
	})

	app.Run(":5000")

}
