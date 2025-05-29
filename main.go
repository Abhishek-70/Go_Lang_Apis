package main

import (
	"restapis/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// it will create require http packages
	r := gin.Default()
	routes.BuildRoutersviaroutes(r)
	r.Run(":4000")   //listen and serve on the port :4000


	//route -> it has been declared in the routes folder

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	//with the same route creating a POST route
	// r.POST("/welcome", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"prop": "goLang",
	// 	})
	// })
	// r.Run(":4030") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
