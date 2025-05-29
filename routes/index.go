package routes

import (
	"restApis/controller"

	"github.com/gin-gonic/gin"
)

func BuildRoutes(r *gin.Engine) {
	r.POST("/create", controller.CreateUser)
	//route
	r.GET("/get", controller.GetUsers)
}
