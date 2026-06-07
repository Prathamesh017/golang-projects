package main

import (
	"fmt"
	"os"
	"pokedex/api"
)
type command struct{
	name string 
	description string
	handlerFunc func(config *api.Config) 
}






func getCommands() []command{
	return [] command{
		{
			name: "help",
			description: "Show this help message",
			handlerFunc: handleHelp,
		},
		{
			name: "exit",
			description: "Exit the REPL",
			handlerFunc: handleExit,
		},
		{
			name: "locations",
			description: "Fetch and display locations from the PokeAPI",
			handlerFunc: handleLocation,
		},
	}
}


func handleExit(_ *api.Config){
	os.Exit(0)
}

func handleHelp(_ *api.Config){
	fmt.Println("Available commands:")
	fmt.Println("help - Show this help message")
	fmt.Println("exit - Exit the REPL")
	fmt.Println("locations - Fetch and display locations from the PokeAPI")
}

func handleLocation(config *api.Config){
	locations,err:=api.GetLocations(config)
	if(err != nil){
		fmt.Printf("Error fetching locations: %v\n", err)
	}

	for _,location:=range locations.Results{
		fmt.Printf("Name: %s\n",location.Name)
	}
}