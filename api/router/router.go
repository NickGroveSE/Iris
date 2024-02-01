package router

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine) {

	app.GET("/music", controllers.GetMusicData)
	app.GET("/health", controllers.GetServerHealth)

}
