package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedex/cache"
)

type Location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Config struct{
	previousUrl *string 
	nextUrl * string
	Cache cache.Cache
}

func GetLocations(config *Config) (Location, error) {
	url := "https://pokeapi.co/api/v2/location"

	if config.nextUrl != nil {
		url = *config.nextUrl
	}

	cachedLocation,ok:=cache.GetCache(config.Cache,url);


	if ok{
		fmt.Println("Cache hit for url: ", url)
		var LocationAreas Location
		err := json.Unmarshal(cachedLocation, &LocationAreas)
		config.nextUrl = LocationAreas.Next
		if err != nil {
			fmt.Printf("Error unmarshalling response: %v\n", err)
			return Location{}, err
		}
		return LocationAreas, nil
	}


	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching locations: %v\n", err)
		return Location{}, err
	}
	if response.StatusCode != 200 {
		return Location{}, fmt.Errorf("invalid status code: %d", response.StatusCode)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return Location{}, err
	}
	var LocationAreas Location
	err = json.Unmarshal(responseData, &LocationAreas)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return Location{}, err
	}
	cache.AddCache(&config.Cache,url,responseData)
	config.nextUrl = LocationAreas.Next
	return LocationAreas, nil
}
