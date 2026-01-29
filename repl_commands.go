package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMapf(cfg *config) error {
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

func commandMapb(cfg *config) error {
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
