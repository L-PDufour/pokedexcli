package main

import (
	"fmt"
	"os"
)

func commandExit(c *config, _ ...string) error {
	fmt.Println("Exiting the Pokedex...")
	os.Exit(0)
	return nil
}
