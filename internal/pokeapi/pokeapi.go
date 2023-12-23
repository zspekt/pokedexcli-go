package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/zspekt/pokedexcli/internal/pokecache"
)

func init() {
	GlobalCache = pokecache.NewCache(10 * time.Second)
	pokeapiClient = NewClient()
}

var pokeapiClient Client

const RootURL string = "https://pokeapi.co/api/v2/"

const EndpointLocArea = "location-area?offset=0&limit=20"

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
		tmp := RootURL + EndpointLocArea
		fullURL = &tmp
	}
	fmt.Println(*fullURL)
	// check if we have this URL's contents cached...
	bytes, err := pokeapiClient.fetchRequest(fullURL)
	if err != nil {
		return LocationAreaResp{}, err
	}
	return unmarshalJson(bytes)
}

func (cl *Client) Explore(cf *Config) (Pkmn, error) {
	var fullURL *string
	var httpResponse *http.Response
	returnVal := Pkmn{}

	httpResponse, err := http.Get(*fullURL)
	if err != nil {
		return Pkmn{}, err
	}
	defer httpResponse.Body.Close()

	body, err := io.ReadAll(httpResponse.Body)
	fmt.Print(body)
	if err != nil {
		return Pkmn{}, err
	}

	return returnVal, err
}

func (r *ExploreAreaResp) listPokemons(c *Config) ([]Pkmn, error) {
	return []Pkmn{}, nil
}

func unmarshalJson(xbyte []byte) (LocationAreaResp, error) {
	r := LocationAreaResp{}

	err := json.Unmarshal(xbyte, &r)
	if err != nil {
		return LocationAreaResp{}, err
	}

	return r, nil
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
