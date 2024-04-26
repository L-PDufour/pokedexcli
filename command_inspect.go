package main

import (
	"errors"
	"fmt"

	"github.com/l-pdufour/pokedexcli/internal/pokeapi"
)

func FormatPokemon(p pokeapi.PokemonResp) string {
	output := fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", p.Name, p.Height, p.Weight)
	for _, stat := range p.Stats {
		output += fmt.Sprintf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	output += "Types:\n"
	for _, t := range p.Types {
		output += fmt.Sprintf("  - Name: %s\n", t.Type.Name)
	}
	return output
}
func commandInspect(c *config, name ...string) error {
	// Ensure there is at least one name argument provided
	if len(name) == 0 {
		return errors.New("missing Pokemon name")
	}

	// Get the name of the Pokemon to inspect
	pokemonName := name[0]

	// Check if the Pokemon is in the caughtPokemonMap
	pokemon, found := caughtPokemonMap[pokemonName]
	if !found {
		return fmt.Errorf("Pokemon %s not found", pokemonName)
	}

	fmt.Println(FormatPokemon(pokemon))

	return nil
}
