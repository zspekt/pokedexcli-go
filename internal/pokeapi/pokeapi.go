// TODO NOW:
// FIX ERRORS:
// 1) NO INPUT ON REPL CRASHES
// 2) CALLING MAPN OR MAPB UNTIL CFG.NEXT/PREV ARE NIL, CRASHES
// 3) TRY TO DRY UP CODE. REPETITION ON ListPrevLocationAreas AND mapb

package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const RootURL string = "https://pokeapi.co/api/v2"

const EndpointLocArea = "/location-area"

func (c *Client) ListLocationAreas(cfg *Config) (LocationAreaResp, error) {
	var response LocationAreaResp

	fullURL := RootURL + EndpointLocArea

	if cfg.NextURL == nil {
		cfg.NextURL = &fullURL
	}

	resp, err := http.Get(*cfg.NextURL)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return LocationAreaResp{}, err
	}

	// for _, location := range response.Results {
	// 	fmt.Println(location.Name)
	// }

	return response, nil
}

func (c *Client) ListPrevLocationAreas(cfg *Config) (LocationAreaResp, error) {
	var response LocationAreaResp

	endpoint := "/location-area"
	fullURL := rootURL + endpoint

	switch {
	case cfg.NextURL == nil:
		break
	case cfg.PreviousURL == nil:
		break
	}

	if cfg.NextURL == nil {
		cfg.NextURL = &fullURL
	}

	resp, err := http.Get(*cfg.PreviousURL)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return LocationAreaResp{}, err
	}

	// for _, location := range response.Results {
	// 	fmt.Println(location.Name)
	// }

	return response, nil
}
