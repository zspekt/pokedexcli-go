package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zspekt/pokedexcli/internal/pokeapi"
)

var cfg = pokeapi.GlobalConfig

func init() {
	cfg = pokeapi.CreateConfig()
}

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
	cfg.Caller = "mapn"

	fmt.Printf(
		"\n\ncfg NEXT --> %v\ncfg PREV --> %v\n\n",
		cfg.NextURL,
		cfg.PreviousURL,
	)

	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListAnyLocationAreas(cfg)
	if err != nil {
		log.Println(err)
	}

	cfg.NextURL = resp.Next
	cfg.PreviousURL = resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	if resp.Previous == nil {
		fmt.Println("*resp.Previous is nil")
	}
	return nil
}

func mapb() error {
	cfg.Caller = "mapb"

	fmt.Printf(
		"\n\ncfg NEXT --> %v\ncfg PREV --> %v\n\n",
		cfg.NextURL,
		cfg.PreviousURL,
	)

	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListAnyLocationAreas(cfg)
	if err != nil {
		fmt.Println(err)
	}

	cfg.NextURL = resp.Next
	cfg.PreviousURL = resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	if resp.Previous == nil {
		fmt.Println("*resp.Previous is nil")
	}
	return nil
}
