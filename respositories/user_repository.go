package respositories

import (
	"restApis/interfaces"
	"restApis/models"
)

type userRepo struct{}

func (r *userRepo) GetUsersById() models.User{
	
	//in real project it's DB call return some values
	return models.User{ID:1,Name: "Alice"}
}

func (r *userRepo) GetAllUsers() []models.User{

	//in real project we interact with DB.

	return []models.User{
		{ID:1,Name: "Alice"},
		{ID:2,Name: "Blice"},
	}
}

func NewUserRepsitory() interfaces.UserRepository{
	return &userRepo{}
}