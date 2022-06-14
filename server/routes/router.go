package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luccasalves/apigo-sample/server/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	main.GET("/", controllers.Greets)

	user := main.Group("users")
	user.GET("/", controllers.GetAll)
	user.GET("/:id", controllers.GetUser)
	user.POST("/", controllers.CreateUser)
	return router
}
