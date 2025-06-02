package controller

import (
	"net/http"
	"restApis/models"
	"restApis/services"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(s services.OrderService) *OrderController{
	return &OrderController{orderService: s}
}

func (c *OrderController) CreateOrder(ctx *gin.Context){
	orderList,_:= c.orderService.CreateOrder(models.Order{})
	ctx.JSON(http.StatusOK,orderList)
}
