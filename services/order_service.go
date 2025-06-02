package services

import (
	"restApis/interfaces"
	"restApis/models"
)

type OrderService interface {
	CreateOrder(models.Order) (orderId int, err error)
}

type orderService struct {
	repo interfaces.IOrder
}

// CreateOrder implements OrderService.
func (o *orderService) CreateOrder(models.Order) (orderId int, err error) {
	return o.repo.CreateOrder(models.Order{})
}

func NewOrderService(r interfaces.IOrder) OrderService {
	return &orderService{repo: r}
}

// type userService struct {
// 	repo interfaces.UserRepository
// }

// // pass the repository dependency
// func NewUserService(r interfaces.UserRepository) UserService {
// 	return &userService{repo: r}
// }
