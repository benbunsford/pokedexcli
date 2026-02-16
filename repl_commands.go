package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
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

func commandCatch(cfg *config, param *string) error {
	if *param == "" {
		fmt.Println("No Pokemon name provided. Please type one after 'catch'!")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", *param)

	pokemonData, err := cfg.pokeapiClient.GetPokemonData(param)
	if err != nil {
		return err
	}
	baseExp := pokemonData.BaseExperience
	catchChance := baseExp*110/100 + 75

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	catchRoll := random.Intn(catchChance)

	if catchRoll >= baseExp {
		fmt.Printf("%s was caught!\n", *param)
		cfg.caughtPokemon[*param] = pokemonData
	} else {
		fmt.Printf("%s escaped!\n", *param)
	}

	return nil
}

func commandInspect (cfg *config, param *string) error {
	if pokemon, ok := cfg.caughtPokemon[*param]; ok {
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" - %v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typ := range pokemon.Types {
			fmt.Printf(" - %v\n", typ.Type.Name)
		}
	} else {
		fmt.Println("You haven't caught that pokemon. Catch it to be able to inspect it.")
	}
	return nil
}

func commandPokedex (cfg *config, param *string) error {
	fmt.Println("Your Pokedex:")

	for key, _ := range cfg.caughtPokemon {
		fmt.Printf(" - %v\n", key)
	}

	return nil
}
