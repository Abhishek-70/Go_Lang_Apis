package controller

import (
	"net/http"
	"restApis/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
    service services.UserService
}

func NewUserController(s services.UserService) *UserController {
    return &UserController{service: s}
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
    users := c.service.GetAllUsers()
    ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetUserById(ctx *gin.Context){
	users:= c.service.GetUserById()
	ctx.JSON(http.StatusOK,users)
}

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func CreateUser(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "mypost",
// 		"abc":     "xyz",
// 	})

// }

// func GetUsers(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "mypost",
// 		"abc":     "xyz",
// 	})

// }

// func DeleteUser(a *gin.Context) {

// }

// func UpdateUser() {

// }
