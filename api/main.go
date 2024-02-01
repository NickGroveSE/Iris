package main

import (
	"api/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	app.Use(cors.Default())

	router.SetupRouter(app)

	app.Run(":5000")

}
