package main

import (
	"restApis/dependency"
	"restApis/routes"
)

func main() {
	// // which will create required http packages
	// r := gin.Default()
	// routes.BuildContainer(r)
	// r.Run(":4000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//Build DI container
	container := dependency.BuildContainer()

	//setup routers
	r := routes.SetupRouter(container.UserController,container.OrderController)
	r.Run(":4002")
}
