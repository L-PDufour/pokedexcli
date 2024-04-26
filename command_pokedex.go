package main

import (
	"errors"
	"fmt"
)

func commandPokedex(c *config, name ...string) error {
	// Ensure there is at least one name argument provided
	if len(name) != 0 {
		return errors.New("too many arguments")
	}

	for _, pokemon := range caughtPokemonMap {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}
