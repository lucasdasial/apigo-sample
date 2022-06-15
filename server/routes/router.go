package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luccasalves/apigo-sample/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	main.GET("/", controllers.Greets)

	user := main.Group("users")
	user.GET("/", controllers.ShowAll)
	user.GET("/:id", controllers.Show)
	user.POST("/", controllers.Create)
	return router
}
