package routes

import (
	"restApis/controller"

	"github.com/gin-gonic/gin"
)

// func BuildRoutes(r *gin.Engine) {
// 	r.POST("/create", controller.CreateUser)
// 	//route
// 	r.GET("/get", controller.GetUsers)
// }

func SetupRouter(userController *controller.UserController) *gin.Engine{
	r:=gin.Default()

	//users Endpoint
	r.GET("/users",userController.GetAllUsers)
	r.GET("/userById",userController.GetUserById)

	return r
}