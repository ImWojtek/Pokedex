package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ImWojtek/Pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			fmt.Println("\nGoodbye!")
			return
		}

		parts := cleanInput(scanner.Text())
		if len(parts) == 0 {
			continue
		}

		cmdName := parts[0]
		args := []string{}
		if len(parts) > 1 {
			args = parts[1:]
		}
		command, exists := getCommands()[cmdName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ReplaceAll(text, ",", " ")
	text = strings.ReplaceAll(text, "/", " ")

	parts := strings.Fields(text)
	for i := range parts {
		parts[i] = strings.ToLower(parts[i])
	}
	return parts
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
			description: "Show available commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List Pokemon at location (ex. explore {location-name})",
			callback:    commandExplore,
		},
	}
}
