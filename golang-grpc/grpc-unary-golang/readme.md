For Doing Grpc , we need first to make a contract between parties invovled , like this will be the request body and this wll be our response.

for that first we define , our `_proto_` file

so now we have define our request , response and service file structure

Once our proto file is define , we need to generate the boilerplate defining this code in the our language 

```
protoc \
--go_out=. \
--go-grpc_out=. \
proto/user.proto
```
This will generate these 2 files
user.pb.go - contains data structures + protobuf serialization/deserialization logic
user_grpc.pb.go - this will have all the gRPC `plumbing/infrastructure`.
this will have all the interfaces, `registartion funcs`  etc.

Let create a grpc server

```go
    lis, err := net.Listen("tcp", ":8080")
	grpcServer:=grpc.NewServer()
	grpcServer.Serve(lis)
```

this is 3 step process , first we have defined 
`net.Listen("tcp", ":8080")` ,
This tells the OS to reserve port 8080. Accept incoming TCP connections.
At this point, we cannot serve any requests yet.This is not a gRPC server.

Next we create a grpc server - just a server object
`grpcServer := grpc.NewServer()`
Create a gRPC engine.
    This object knows how to:
    Handle HTTP/2
    Understand protobuf payloads
    Route RPC method calls
    Support streaming
At this point, it's still idle.

In the last step
`grpcServer.Serve(lis)`
    * Connect the listener (lis) with the gRPC server.
    * Start accepting connections on port 8080.
    * The server is now running.


At this point , we don't have routes or anything to serve the actual request,Now we have to actually write the proto contract implemention for serveer

First we create a struct , this struct essentially will be mapped to the _proto_ golang file create to make sure we are defining all the services in correct format for satisfying the contract

```
type UserService struct{
	   pb.UnimplementedUserServiceServer
}
```

then we actual have to register , this will map to the files we generated making sure all the rpc functions are made in contract are correctly defined or not
```
    pb.RegisterUserServiceServer(
    grpcServer,&UserService{},
	)
```

and then we have to implement our rpc functions similar to contract to satisfy the struct

```
func (s *UserService) GetUser (ctx context.Context,req *pb.UserRequest) (res *pb.UserResponse,err error){
   return &pb.UserResponse{
	Name: "Prathamesh",
	Id: 1,
   },nil
}
```


Let understand this `pb.UnimplementedUserServiceServer` . even without defining this , essentially it would work fine

so the problem appears ,if tommorow you update you proto file and regenerate new methods , but then go will start complainign you don't implementaion yet 

this creates a default implementation for you to avoid go errors, something like 
```
func (UnimplementedUserServiceServer) GetUser(...) {
   return "not implemented"
}
```
so you don't get errors


Defining the Client

JS Client  defining is actually simpler , it is simply htting the function running on port 8081, also interesting we have is js as it is dynamic language , don't use types , so don't have boilerplate interface code as it reads the file when we run the program directly doesn't need type safety beforehand and is flexible

even we can do sme thing even we are using ts , but generally people using it cause it gets all the type safety feature , else no point in using typescript over js in first place