package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config, param *string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, param *string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMapf(cfg *config, param *string) error {
	locationBatch, err := cfg.pokeapiClient.GetMapData(cfg.nextLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = locationBatch.Next
	cfg.previousLocationUrl = locationBatch.Previous

	for _, location := range locationBatch.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config, param *string) error {
	if cfg.previousLocationUrl == nil {
		return errors.New("you're on the first page")
	}

	locationBatch, err := cfg.pokeapiClient.GetMapData(cfg.previousLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = locationBatch.Next
	cfg.previousLocationUrl = locationBatch.Previous

	for _, location := range locationBatch.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExplore(cfg *config, param *string) error {
	if *param == "" {
		fmt.Println("No location-area name provided. Please type one after 'explore'!")
	}

	locationData, err := cfg.pokeapiClient.GetLocationData(param)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", *param)
	fmt.Println("Found Pokemon:")

	for _, encounter := range locationData.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
