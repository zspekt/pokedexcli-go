package pokeapi

import (
	"net/http"
	"time"
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

type Area struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Config struct {
	NextURL     *string
	PreviousURL *string
}

func CreateConfig() *Config {
	// assign the url to something so we can pass the pointer
	// we want a default value for something something not crash when pointer is nil
	// i kinda forgot why i was doing this in the first place
	nextURL := RootURL + EndpointLocArea
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