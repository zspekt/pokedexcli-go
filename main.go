// TODO NOW:
// FIX ERRORS:
// 1) NO INPUT ON REPL CRASHES
// 2) CALLING MAPN OR MAPB UNTIL CFG.NEXT/PREV ARE NIL, CRASHES
// 3) TRY TO DRY UP CODE. REPETITION ON ListPrevLocationAreas AND mapb

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
