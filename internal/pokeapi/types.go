package pokeapi

import (
	"net/http"
	"time"

	"github.com/zspekt/pokedexcli/internal/pokecache"
)

var (
	GlobalConfig *Config
	GlobalCache  *pokecache.Cache
)

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

type Config struct {
	AreaToExplore *string
	NextURL       *string
	PreviousURL   *string
	// caller would allow to change behaviour depending on
	// whether mapn or mapb is calling ListLocationAreas
	Caller string
}

func CreateConfig() *Config {
	// assign the url to something so we can pass the pointer
	// we want a default value for something something not crash when pointer is nil
	// i kinda forgot why i was doing this in the first place
	nextURL := RootURL + EndpointLocAreaListing
	return &Config{
		NextURL: &nextURL,
	}
}

type LocationAreaResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []Area  `json:"results"`
}

type Area struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ExploreAreaResp struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon        Pokemon `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
