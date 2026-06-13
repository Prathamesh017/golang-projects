package routes 

import (
	"github.com/gin-gonic/gin"
	"restaurent-managment/handler"
)

func RegisterAuthGroup(router *gin.RouterGroup,authHandler *handler.AuthHandler,
) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", authHandler.SignUp)
		authGroup.POST("/login",authHandler.Login)
	}
}