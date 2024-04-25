package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const NumberOfRequest int = 20

const green string = "\033[32m"
const reset string = "\033[0m"

func (c *Client) ListLocationArea(PageURL *string) (PokemonLocationArea, error) {
	endpoint := "/location-area"

	var fullURL string
	if PageURL != nil {
		fullURL = *PageURL
	} else {
		fullURL = baseURL + endpoint
	}

	if cachedRespBytes, ok := c.cache.Get(fullURL); ok {
		var cachedResp PokemonLocationArea
		if err := json.Unmarshal(cachedRespBytes, &cachedResp); err != nil {
			return PokemonLocationArea{}, fmt.Errorf("failed to unmarshal cached response: %v", err)
		}
		return cachedResp, nil
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonLocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return PokemonLocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonLocationArea{}, err
	}

	locationAreaResp := PokemonLocationArea{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return PokemonLocationArea{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreaResp, nil
}
