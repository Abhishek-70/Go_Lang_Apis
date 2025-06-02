package services

import (
	"restApis/interfaces"
	"restApis/models"
)

type UserService interface {
	//her we have to define the methods/func from interface package having some defined logic
	GetAllUsers() []models.User

	GetUserById() models.User
}

type userService struct {
	repo interfaces.UserRepository
}

//declare all the method implementation from the repository

func (s *userService) GetAllUsers() []models.User {
	return s.repo.GetAllUsers()
}

//if we need to access another GetUSerById method we need to define here

func (s *userService) GetUserById() models.User{
	return s.repo.GetUsersById()
}

// pass the repository dependency
func NewUserService(r interfaces.UserRepository) UserService {
	return &userService{repo: r}
}