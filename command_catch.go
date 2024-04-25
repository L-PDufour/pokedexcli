package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func commandCatch(c *config, name ...string) error {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	if len(name) == 0 || name[0] == "" {
		fmt.Println("Need a pokemon name")
		return nil
	}

	// Simulate a random chance of catching the Pokemon
	chance := rand.Float64() // Generate a random float between 0 and 1
	if chance < 0.5 {        // Adjust the threshold to change the catch rate
		fmt.Println("Failed to catch", name[0])
		return nil
	}

	// If the catch is successful, store the Pokemon in a map
	resp, err := c.pokeapiClient.CatchPokemon(name[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Caught", resp.Name)
	pokemon := Pokemon{Name: resp.Name}
	storeInMap(resp.Name, pokemon) // Call a function to store the caught Pokemon in a map

	return nil
}
