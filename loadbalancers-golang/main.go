package main 

import (
	"fmt"
	"net/url"
	"net/http"
)


type Server struct {
	url string
}

type LoadBalancer struct {
	port string //it's own port
	servers []Server
	robinRoundCount int
}

func createServer(addRes string) Server{
	redirectUrl,err:=url.Parse(addRes);
	if err!=nil{
		panic(err);
	}
	return  Server{
		url:redirectUrl.String(),
	}
}


func (s Server) isAlive() bool{
	//hardcoding is fine for testing
	return true
}

func ( s Server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Request is served by %s", s.url);
	//In real world, we would forward the request to the actual server and return the response back to the client
	
}


func (lb *LoadBalancer) getNextServer() *Server{
	latestCheck:=lb.robinRoundCount%len(lb.servers);
	fmt.Println(latestCheck);
	server:=lb.servers[latestCheck];
	if(server.isAlive()){
		lb.robinRoundCount++;
		return &server;
	}
	return lb.getNextServer()
}


func redirect(w http.ResponseWriter, r *http.Request,lb *LoadBalancer){
	server:=lb.getNextServer();
	server.ServeHTTP(w,r);
}


func main(){
	server1:=createServer("http://reddit.com")
	server2:=createServer("http://google.com")
	server3:=createServer("http://facebook.com")

	lb:=LoadBalancer{
		port:"8080",
		servers: []Server{server1, server2, server3},
		robinRoundCount: 0,
	}

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		redirect(w,r,&lb)
	});
	fmt.Println("Load Balancer is running on port", lb.port);
	err:=http.ListenAndServe(":"+lb.port,nil);
	if(err!=nil){
		panic(err)
	}
	
}