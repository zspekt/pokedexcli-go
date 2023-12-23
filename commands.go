package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zspekt/pokedexcli/internal/pokeapi"
)

var Cfg = pokeapi.GlobalConfig

func init() {
	Cfg = pokeapi.CreateConfig()
}

func commandExit(*pokeapi.Config) error {
	os.Exit(0)
	return nil
}

func commandHelp(*pokeapi.Config) error {
	fmt.Printf("\nThis is the Go Pokedex CLI. Version 0.0.0\n\n")
	for _, value := range getCommand() {
		fmt.Printf("%v: %v\n", value.name, value.description)
	}
	fmt.Println()
	return nil
}

func mapn(*pokeapi.Config) error {
	Cfg.Caller = "mapn"

	fmt.Printf(
		"\n\nCfg NEXT --> %v\nCfg PREV --> %v\n\n",
		Cfg.NextURL,
		Cfg.PreviousURL,
	)

	resp, err := pokeapi.ListAnyLocationAreas(Cfg)
	if err != nil {
		log.Println(err)
	}

	Cfg.NextURL = resp.Next
	Cfg.PreviousURL = resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	if resp.Previous == nil {
		fmt.Println("*resp.Previous is nil")
	}
	return nil
}

func mapb(*pokeapi.Config) error {
	Cfg.Caller = "mapb"

	fmt.Printf(
		"\n\nCfg NEXT --> %v\nCfg PREV --> %v\n\n",
		Cfg.NextURL,
		Cfg.PreviousURL,
	)

	resp, err := pokeapi.ListAnyLocationAreas(Cfg)
	if err != nil {
		fmt.Println(err)
	}

	Cfg.NextURL = resp.Next
	Cfg.PreviousURL = resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	if resp.Previous == nil {
		fmt.Println("*resp.Previous is nil")
	}
	return nil
}

func explore(c *pokeapi.Config) error {
	c.AreaToExplore = &CmdWritten[1]

	ExploreAreaResp, err := pokeapi.Explore(c)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Let's see what's in ", c.AreaToExplore, "...")
	for _, p := range ExploreAreaResp.PokemonEncounters {
		fmt.Println("\t\t--> ", p.Pokemon.Name)
	}

	return nil
}
