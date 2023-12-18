package main

import (
	"github.com/zspekt/pokedexcli/internal/pokeapi"
)

var cfg *pokeapi.Config

func main() {
	// pokeapiClient := pokeapi.NewClient()
	//
	// resp, err := pokeapiClient.ListLocationAreas()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp)
	repl()
}
