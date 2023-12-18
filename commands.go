package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zspekt/pokedexcli/internal/pokeapi"
)

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf("\nThis is the Go Pokedex CLI. Version 0.0.0\n\n")
	for _, value := range getCommand() {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	fmt.Println()
	return nil
}

func mapn() error {
	// if cfg doesn't exist, we create it. runtime error if we don't
	if cfg == nil {
		cfg = pokeapi.CreateConfig()
	}

	fmt.Printf("\n\ncfg NEXT --> %v\ncfg PREV --> %v\n\n", cfg.NextURL, cfg.PreviousURL)

	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas(cfg)
	if err != nil {
		log.Fatal(err)
	}

	cfg.NextURL = resp.Next
	cfg.PreviousURL = resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	if resp.Previous == nil {
		fmt.Println("*resp.Previous is nil")
	}
	fmt.Printf("\n\nAFTER http get. Next: %v\nPrevious: %v\n", resp.Next, resp.Previous)
	return nil
}

func mapb() error {
	// if cfg doesn't exist, we create it. runtime error if we don't
	if cfg == nil {
		cfg = pokeapi.CreateConfig()
	}

	fmt.Printf("\n\ncfg NEXT --> %v\ncfg PREV --> %v\n\n", cfg.NextURL, cfg.PreviousURL)

	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListPrevLocationAreas(cfg)
	if err != nil {
		log.Fatal(err)
	}

	cfg.NextURL = resp.Next
	cfg.PreviousURL = resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	if resp.Previous == nil {
		fmt.Println("*resp.Previous is nil")
	}
	fmt.Printf("\n\nNext: %v\nPrevious: %v", *resp.Next, resp.Previous)
	return nil
}
