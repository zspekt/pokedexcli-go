package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const RootURL string = "https://pokeapi.co/api/v2"

const EndpointLocArea = "/location-area"

func (c *Client) ListAnyLocationAreas(cfg *Config) (LocationAreaResp, error) {
	var response LocationAreaResp
	var fullURL *string

	// figure out which URL we want based on who called
	switch {
	case cfg.Caller == "mapn":
		fullURL = cfg.NextURL
	case cfg.Caller == "mapb":
		fullURL = cfg.PreviousURL
	}

	// if URL pointer is nil, figure out which error message we should return
	// there is definitely a cleaner way to do this
	switch {
	// if both URLs are nil, reset to beginning
	case cfg.NextURL == nil && cfg.PreviousURL == nil:
		tmp := RootURL + EndpointLocArea
		fullURL = &tmp
	case cfg.Caller == "mapn" && fullURL == nil:
		return LocationAreaResp{}, errors.New("Reached the end of the line, bud")
	case cfg.Caller == "mapb" && fullURL == nil:
		return LocationAreaResp{}, errors.New("Nowhere to go back to, buddy")
	}

	resp, err := http.Get(*fullURL)
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

//
// func (c *Client) ListLocationAreas(cfg *Config) (LocationAreaResp, error) {
// 	var response LocationAreaResp
//
// 	fullURL := RootURL + EndpointLocArea
//
// 	if cfg.NextURL == nil {
// 		cfg.NextURL = &fullURL
// 	}
//
// 	resp, err := http.Get(*cfg.NextURL)
// 	if err != nil {
// 		return LocationAreaResp{}, err
// 	}
// 	defer resp.Body.Close()
//
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return LocationAreaResp{}, err
// 	}
//
// 	err = json.Unmarshal(body, &response)
// 	if err != nil {
// 		return LocationAreaResp{}, err
// 	}
//
// 	// for _, location := range response.Results {
// 	// 	fmt.Println(location.Name)
// 	// }
//
// 	return response, nil
// }
//
// func (c *Client) ListPrevLocationAreas(cfg *Config) (LocationAreaResp, error) {
// 	var response LocationAreaResp
//
// 	if cfg.PreviousURL == nil {
// 		return LocationAreaResp{}, errors.New("Nowhere to go back to.")
// 	}
//
// 	resp, err := http.Get(*cfg.PreviousURL)
// 	if err != nil {
// 		return LocationAreaResp{}, err
// 	}
// 	defer resp.Body.Close()
//
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return LocationAreaResp{}, err
// 	}
//
// 	err = json.Unmarshal(body, &response)
// 	if err != nil {
// 		return LocationAreaResp{}, err
// 	}
//
// 	// for _, location := range response.Results {
// 	// 	fmt.Println(location.Name)
// 	// }
//
// 	return response, nil
// }
