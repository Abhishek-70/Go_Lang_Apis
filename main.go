package main

import (
	"restApis/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// which will create required http packages
	r := gin.Default()
	routes.BuildRoutes(r)
	r.Run(":4000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
