package routes

import (
	"github.com/gin-gonic/gin"
	"restaurent-managment/handler"
)


func RegisterUserRoutes(router *gin.RouterGroup,userHandler *handler.UserHandler,){
    userGroup := router.Group("/user")
	{
		userGroup.GET("/", userHandler.GetUser)
	}
}