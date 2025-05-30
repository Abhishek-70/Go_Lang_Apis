package interfaces

import "restApis/models"

type UserRepository interface {
	GetAllUsers() []models.User
	GetUsersById() models.User
}