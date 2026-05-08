package main

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
)
type Pokemon struct {
	Name        string `json:"name"`
	PokeMonType string `json:"type"`
	XP          int    `json:"xp"`
	Power       string `json:"power"`
	Level       int    `json:"level"`
}


func initRedisClient() (*redis.Client, context.Context) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})

	ctx := context.Background()
	return rdb, ctx
}

func getAllTypePokemon(rdb *redis.Client, ctx context.Context, pokemonType string) ([]Pokemon, error) {
	var val string="pokemon:"+pokemonType+":"+"*"
	keys, err := rdb.Keys(ctx, val).Result()
	if err != nil {
		return nil, err
	}

	var pokemons []Pokemon=make([]Pokemon,0)
	for _,val := range keys{
		var p Pokemon
		var pokemon=rdb.Get(ctx,val).Val()
		err := json.Unmarshal([]byte(pokemon), &p)
		if err != nil {
			return nil, err
		}
		pokemons=append(pokemons,p)
	}
	return pokemons, nil
}

func makePokemonHandler(rdb *redis.Client, ctx context.Context, pokemonType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pokemons, err := getAllTypePokemon(rdb, ctx, pokemonType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(pokemons)
	}
}

func main() {
	rdb, ctx := initRedisClient()
	
	http.HandleFunc("/water", makePokemonHandler(rdb, ctx, "water"))
	http.HandleFunc("/fire", makePokemonHandler(rdb, ctx, "fire"))
	http.HandleFunc("/grass", makePokemonHandler(rdb, ctx, "grass"))
	http.HandleFunc("/electric", makePokemonHandler(rdb, ctx, "electric"))
	http.HandleFunc("/psychic", makePokemonHandler(rdb, ctx, "psychic"))
	http.HandleFunc("/bug", makePokemonHandler(rdb, ctx, "bug"))
	http.HandleFunc("/rock", makePokemonHandler(rdb, ctx, "rock"))
	http.HandleFunc("/ghost", makePokemonHandler(rdb, ctx, "ghost"))
	http.HandleFunc("/dragon", makePokemonHandler(rdb, ctx, "dragon"))
	http.HandleFunc("/dark", makePokemonHandler(rdb, ctx, "dark"))
	http.HandleFunc("/steel", makePokemonHandler(rdb, ctx, "steel"))
	http.HandleFunc("/fairy", makePokemonHandler(rdb, ctx, "fairy"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}