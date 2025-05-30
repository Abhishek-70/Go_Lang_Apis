package dependency

import (
	"restApis/controller"
	"restApis/services"
	"restApis/respositories"
)

type Container struct {
	UserController *controller.UserController
}

func BuildContainer() *Container {
	//1.Models->Create Interface->Repository->Service->Controller

	// Wiring dependencies manually
	userRepo := respositories.NewUserRepsitory()
	userService := services.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	return &Container{
		UserController: userController,
	}
}
