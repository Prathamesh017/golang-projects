package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"restaurent-managment/routes"
	"restaurent-managment/db"
	"context"
	"restaurent-managment/services"
	"restaurent-managment/handler"
)

func main(){
	conn,err2:=db.ConnectDB()
	userService:=services.NewUserService(conn);
	authService:=services.NewAuthService(conn,*userService);
	authHandler:=handler.NewAuthHandler(authService)
	userHandler:=handler.NewUserHandler(userService)
	

	if err2 !=nil{
			panic(err2)
	}

	defer conn.Close(context.Background())

	router:=gin.Default()
	api:=router.Group("/api")
	routes.RegisterAuthGroup(api,authHandler)
	routes.RegisterUserRoutes(api,userHandler)
	router.GET("/health", func(context *gin.Context){

		context.JSON(200,gin.H{
			"message":"Healthy",
		})
	})

	fmt.Println("Server is running on port 8080")
	err:=router.Run(":8080")
	if err!=nil{
		fmt.Println("Error starting server:", err)
	}

}