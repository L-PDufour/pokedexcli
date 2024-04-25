package main

import (
	"fmt"
	"log"
)

func commandExplore(c *config, name ...string) error {

	resp, err := c.pokeapiClient.ListPokemonArea(name[0])
	if err != nil {
		log.Fatal(err)
	}
	for _, area := range resp.PokemonEncounters {
		fmt.Println(area.Pokemon.Name)
	}

	return nil
}
