package main

import (
	"time"

	"github.com/l-pdufour/pokedexcli/internal/pokeapi"
)

var caughtPokemonMap = make(map[string]pokeapi.PokemonResp)

func storeInMap(name string, pokemon pokeapi.PokemonResp) {
	caughtPokemonMap[name] = pokemon
}

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	readline(cfg)
}
