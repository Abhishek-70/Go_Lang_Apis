package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
     "messsage":"Hello",
	 "regarding":"controllerLayer",
	})
}

func GetUser(a *gin.Context){
	a.JSON(http.StatusOK,gin.H{
		"user":"abhishek",
		"details":"sde",
	})
}

func DeleteUser(){
 
}

func UpdateUser(c *gin.Context){

}