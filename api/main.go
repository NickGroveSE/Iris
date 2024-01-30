package main

import (
	"api/router"
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

	router.SetupRouter(app)

	app.Run(":5000")

}
