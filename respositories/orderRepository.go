package respositories

import (
	"restApis/interfaces"
	"restApis/models"
)

type orderRepo struct{}

// func (r *userRepo) GetUsersById() models.User{

// 	//in real project it's DB call return some values
// 	return models.User{ID:1,Name: "Alice"}
// }

func (r *orderRepo) CreateOrder(models.Order) (orderId int,err error){
	p1:=models.Order{
		//OrderId: 12,Type: "fruit",Price: 120,Quantity: 4,
	}
	if p1.OrderId !=0{
  return p1.OrderId,nil
	}else{
		return 0,err
	}
}

func NewOrderRepository() interfaces.IOrder{
return &orderRepo{}
}
