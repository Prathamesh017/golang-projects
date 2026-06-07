package main 

import (
	"fmt"
	"net/http";
	"github.com/gorilla/mux"
	"encoding/json"
)

type Movie struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Rating int `json:"rating"`
	Director *Director `json:"director"`
}

type Director struct{
	FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
}

type MovieResponse struct{
	Message string `json:"message"`
	Movie  []Movie `json:"movie"`
}

var movies []Movie

func homeRoute(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w,"Welcome to the Home Page")
}

func getMovies(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovieByID(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	for _, movie := range movies {
		if movie.ID == id{
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.NotFound(w, r)
}

func getMovieByName(w http.ResponseWriter , r *http.Request){
	vars:=mux.Vars(r);
	name:=vars["name"];
	fmt.Println("Request is Being Hit")
	for _,movie :=range movies{
		if movie.Name== name{
			json.NewEncoder(w).Encode(movie)
			w.Header().Set("Content-Type", "application/json")
			return
		}
	}
	http.NotFound(w, r)
}

func createMovie(w http.ResponseWriter , r *http.Request){
    var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
  }
   movie.ID=fmt.Sprintf("%d",len(movies)+1)
   movies=append(movies,movie)
   w.Header().Set("Content-Type", "application/json")
   json.NewEncoder(w).Encode(
	MovieResponse{
	Message:"Movie created successfully",
	Movie:movies,
})
}

func initMovies() [] Movie {
    director := Director{
        FirstName: "Christopher",
        LastName: "Nolan",
    }

    movies:= []Movie{
        {
            ID: "1",
            Name: "Inception",
            Rating: 9,
            Director: &director,
        },
        {
            ID: "2",
            Name: "Interstellar",
            Rating: 10,
            Director: &director,
        },
    }

	return movies;
}



func main(){
	router:=mux.NewRouter();
	movies = initMovies()	
	router.HandleFunc("/",homeRoute).Methods("GET")
	router.HandleFunc("/movies",getMovies).Methods("GET")
	router.HandleFunc("/movies/id/{id}",getMovieByID).Methods("GET")
	router.HandleFunc("/movies/name/{name}",getMovieByName).Methods("GET")

	router.HandleFunc("/movies",createMovie).Methods("POST")
	fmt.Println("Server is running on port 8080")
	err:=http.ListenAndServe(":8080",router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}