package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/zspekt/pokedexcli/internal/pokecache"
)

var (
	pokeapiClient  Client
	CaughtPokemons map[string]Pokemon
)

const (
	RootURL                  = "https://pokeapi.co/api/v2/"
	EndpointLocAreaExploring = "location-area/"
	EndpointLocAreaListing   = "location-area?offset=0&limit=20"
	EndpointPokemon          = "pokemon/"
)

func init() {
	GlobalCache = pokecache.NewCache(5 * time.Minute)
	pokeapiClient = NewClient()
	CaughtPokemons = map[string]Pokemon{}
}

// makes the http request and unmarshals the data. the reference to a config struct
// that is passed to it provides the URLs (which it will overwrite), and    also
// tell the function who called, so it knows which URL to get
func ListAnyLocationAreas(cfg *Config) (LocationAreaResp, error) {
	var fullURL *string

	// figure out which URL we want based on who called
	switch cfg.Caller {
	case "mapn":
		fmt.Println("caller is mapn and this is NEXTURL --> ", cfg.NextURL)
		fullURL = cfg.NextURL
	case "mapb":
		fmt.Println("caller is mapb and this is PREVURL --> ", cfg.PreviousURL)
		fullURL = cfg.PreviousURL
	}

	if fullURL == nil || *fullURL == "" {
		fmt.Println("URL pointer is nil or empty. Initializing with default value.")
		tmp := RootURL + EndpointLocAreaListing
		fullURL = &tmp
	}
	fmt.Println(*fullURL)
	// check if we have this URL's contents cached...
	bytes, err := pokeapiClient.fetchRequest(fullURL)
	if err != nil {
		return LocationAreaResp{}, err
	}
	return unmarshalJson[LocationAreaResp](bytes)
}

func unmarshalJson[T any](xbyte []byte) (T, error) {
	var returnVal T

	err := json.Unmarshal(xbyte, &returnVal)
	if err != nil {
		log.Fatal(err)
		return returnVal, err
	}

	return returnVal, nil
}

// retrieves from cache and returns or makes HTTP request and adds it to the cache
func (c *Client) fetchRequest(url *string) ([]byte, error) {
	// return value
	var bytes []byte
	var httpResponse *http.Response

	if bytes, ok := GlobalCache.Get(*url); ok {
		fmt.Printf("\n\n%v\n\n", "USING CACHE")
		return bytes, nil
	}
	fmt.Printf("\n\n%v\n\n", "NOT USING CACHE")

	httpResponse, err := http.Get(*url)
	if err != nil {
		return []byte{}, err
	}
	defer httpResponse.Body.Close()

	bytes, err = io.ReadAll(httpResponse.Body)
	if err != nil {
		return []byte{}, err
	}
	// adding the entry to the cache
	GlobalCache.Add(*url, bytes)

	return bytes, nil
}

func Explore(c *Config) (ExploreAreaResp, error) {
	url := RootURL + EndpointLocAreaExploring + *c.Argument
	// url := test
	returnVal := ExploreAreaResp{}

	fmt.Println("\n\n", url, "\n\n")
	bytes, err := pokeapiClient.fetchRequest(&url)
	if err != nil {
		log.Fatal(err)
		return ExploreAreaResp{}, err
	}

	returnVal, err = unmarshalJson[ExploreAreaResp](bytes)
	if err != nil {
		log.Fatal(err)
		return ExploreAreaResp{}, err
	}

	return returnVal, nil
}

func Catch(c *Config) error {
	//
	var (
		baseXp            int
		baseChanceToCatch float32 = 0.50
		baseXpWeight      float32 = 0.05
		modifier          float32
		modifiedChance    float32 = baseChanceToCatch - modifier
		pokemonResp       PokemonResp
		pokemonToCatch    string = *c.Argument
	)

	url := RootURL + EndpointPokemon + pokemonToCatch

	bytes, err := pokeapiClient.fetchRequest(&url)
	if err != nil {
		log.Println(err)
		return err
	}

	pokemonResp, err = unmarshalJson[PokemonResp](bytes)
	if err != nil {
		log.Println(err)
		return err
	}

	baseXp = pokemonResp.BaseExperience
	modifier = float32(baseXp) * baseXpWeight
	rand := rand.Float32()

	fmt.Printf("\nThrowing pokeball at %v ...\n", pokemonToCatch)
	if rand < modifiedChance {
		fmt.Printf("\nSuccess! You've caught %v.\n", pokemonToCatch)
		CaughtPokemons[pokemonToCatch] = Pokemon{
			Name: pokemonToCatch,
			URL:  url,
		}
		return nil
	} else {
		fmt.Printf("\nTrying to catch %v has failed.\n", pokemonToCatch)
		return nil
	}
}
