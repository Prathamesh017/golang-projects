package handler

import (
	"fmt"
	"restaurent-managment/services"
	"github.com/gin-gonic/gin"
	"restaurent-managment/types"
)
type AuthHandler struct {
	authService *services.AuthService

}



func NewAuthHandler(auth *services.AuthService) *AuthHandler{
	return &AuthHandler{
		authService:auth,
	}
}

//Context is combination of request and response, something like (req,res) in express
//it provides methods to read request data and write response data
//if we check signature , it has both the responseWriter and request as part of it, so we can read and write data using this context
func (handler *AuthHandler) SignUp( context *gin.Context) {

//Defining a object of type SignupRequest to bind the incoming json data to it
var signupRequest types.SignupRequest;
err:=context.ShouldBindJSON(&signupRequest);

if err!=nil{
	message:=fmt.Sprintf("Invalid request data: %v", err)
	context.JSON(400, gin.H{
		"message": message,
	})
	return
}

res,err:=handler.authService.SignUp(signupRequest)

if err!=nil{
	message:=fmt.Sprintf("Error signing up user: %v", err)
	context.JSON(500, gin.H{
		"message": message,
	})
	return

}
context.JSON(201, res)
}


func (handler *AuthHandler) Login(context *gin.Context){

	var loginRequest types.LoginRequest

	err:=context.ShouldBindJSON(&loginRequest);

	if err!=nil{
		message:=fmt.Sprintf("Invalid Request Body:%v",err);
		context.JSON(400,gin.H{
			"message":message,
		})
	}

	res,err:=handler.authService.Login(loginRequest.Email,loginRequest.Password)

	if err!=nil{
	message:=fmt.Sprintf("Error signing up user: %v", err)
	context.JSON(500, gin.H{
		"message": message,
	})
	return

}
		context.JSON(201, res)

}