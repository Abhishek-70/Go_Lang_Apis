package interfaces


import (
    "restApis/models"
)


type IOrder interface {
    CreateOrder(models.Order) (orderId int, err error)
}