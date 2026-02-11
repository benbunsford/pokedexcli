package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/benbunsford/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationUrl     *string
	previousLocationUrl *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) == 0 {
			continue
		}

		param := ""
		if len(cleanedInput) > 1 {
			param = cleanedInput[1]
		}

		command, ok := getCommands()[cleanedInput[0]]
		if ok {
			err := command.callback(cfg, &param)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	cleaned := strings.Fields(lower)
	return cleaned
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location-areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location-areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays the pokemon that reside in a location-area.",
			callback:    commandExplore,
		},
	}
}
