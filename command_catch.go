package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func commandCatch(c *config, name ...string) error {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	if len(name) == 0 || name[0] == "" {
		fmt.Println("Need a pokemon name")
		return nil
	}

	chance := rng.Float64()
	if chance < 0.5 {
		fmt.Println("Failed to catch", name[0])
		return nil
	}

	// If the catch is successful, store the Pokemon in a map
	resp, err := c.pokeapiClient.CatchPokemon(name[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Caught", resp.Name)
	storeInMap(resp.Name, resp) // Call a function to store the caught Pokemon in a map

	return nil
}
