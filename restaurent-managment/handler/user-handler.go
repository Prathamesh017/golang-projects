package handler

import (
	"restaurent-managment/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{
	userService *services.UserService
}

func NewUserHandler(us *services.UserService) *UserHandler{
	return &UserHandler{
		userService: us,
	}
}

func (user *UserHandler) GetUser(context *gin.Context){
	 context.JSON(200,gin.H{
		"message":"User Found",
	})
}