package main

import (
	"time"

	"github.com/l-pdufour/pokedexcli/internal/pokeapi"
)

type Pokemon struct {
	Name string // Assuming there's a Name field in your Pokemon struct
	// Add other fields as needed
}

var caughtPokemonMap = make(map[string]Pokemon)

func storeInMap(name string, pokemon Pokemon) {
	caughtPokemonMap[name] = pokemon
}

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	readline(cfg)
}
