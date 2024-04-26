package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/l-pdufour/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextURL       *string
	prevURL       *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"nmap": {
			name:        "nmap",
			description: "Next area locations",
			callback:    commandNextLocationAreas,
		},
		"pmap": {
			name:        "pmap",
			description: "Previous aera locations",
			callback:    commandPrevLocationAreas,
		},
		"explore": {
			name:        "explore <area>",
			description: "List of pokemons",
			callback:    commandExplore,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "List of pokemons",
			callback:    commandInspect,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "List of pokemons",
			callback:    commandCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List of caught pokemons",
			callback:    commandPokedex,
		}}
}

func cleanInput(str string) []string {
	input_lowered := strings.ToLower(str)
	input_tokens := strings.Fields(input_lowered)

	return input_tokens
}

func readline(c *config) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)     // Trim whitespace
		inputTokens := strings.Fields(input) // Split input into tokens

		if len(inputTokens) == 0 {
			continue
		}

		commandName := inputTokens[0]
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Invalid command. Type 'help' for available commands.")
			continue
		}

		var parameters []string
		if len(inputTokens) > 1 {
			parameters = inputTokens[1:]
		}

		err := command.callback(c, parameters...)
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
		}
	}
}
