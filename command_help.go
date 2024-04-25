package main

import "fmt"

func commandHelp(c *config, _ ...string) error {
	fmt.Println("Help message:")
	commandList := getCommands()
	for _, cmd := range commandList {
		fmt.Printf(" - %s : %s\n", cmd.name, cmd.description)
	}
	return nil
}
