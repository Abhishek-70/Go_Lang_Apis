package routes

import (
	"restapis/controllers"

	"github.com/gin-gonic/gin"
)

func BuildRoutersviaroutes(r *gin.Engine){

	r.POST("/create",controllers.CreateUser)

	//1.taken care in the controller layer
	// r.POST("/welcome", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "PostThisMessage",
	// 	})
	// })

	//routes
 r.GET("/getUserDetails",controllers.GetUser)

	//2.It also created in the controller layer.
	//route to different path
	// r.GET("/value", func(c1 *gin.Context){
	// 	c1.JSON(http.StatusOK,gin.H{
	// 		"name":"Abhishek",
	// 		"address":"DLW",
	// 	})
	// 	})
}
