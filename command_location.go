package main

import (
	"fmt"
	"log"
)

// TODO define max pages to not got overit
func commandNextLocationAreas(c *config, _ ...string) error {

	resp, err := c.pokeapiClient.ListLocationArea(c.nextURL)
	if err != nil {
		log.Fatal(err)
	}
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	c.nextURL = resp.Next
	c.prevURL = resp.Previous
	return nil
}

func commandPrevLocationAreas(c *config, _ ...string) error {

	resp, err := c.pokeapiClient.ListLocationArea(c.prevURL)
	if err != nil {
		log.Fatal(err)
	}
	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	c.nextURL = resp.Next
	c.prevURL = resp.Previous

	return nil
}
