package main

import (
	"context"
	"fmt"
	"net"
	"google.golang.org/grpc"
	pb "grpc-unary-golang/proto"
)


type UserService struct{
	   pb.UnimplementedUserServiceServer
}

func (s *UserService) GetUser (ctx context.Context,req *pb.UserRequest) (res *pb.UserResponse,err error){
   return &pb.UserResponse{
	Name: "Prathamesh",
	Id: 1,
   },nil
}
func main(){
	lis, err := net.Listen("tcp", ":8081")
	grpcServer:=grpc.NewServer()
	pb.RegisterUserServiceServer(
    grpcServer,&UserService{},
	)
	fmt.Println("Listening on PORT")
	grpcServer.Serve(lis)
	
	if err !=nil{
		fmt.Println("Errors Poping UP")
	}
	
}
