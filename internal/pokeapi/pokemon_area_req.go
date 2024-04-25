package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemonArea(name string) (PokemonAreaResp, error) {
	endpoint := "/location-area/"

	fullURL := baseURL + endpoint + name

	if cachedRespBytes, ok := c.cache.Get(fullURL); ok {
		var cachedResp PokemonAreaResp
		if err := json.Unmarshal(cachedRespBytes, &cachedResp); err != nil {
			return PokemonAreaResp{}, fmt.Errorf("failed to unmarshal cached response: %v", err)
		}
		return cachedResp, nil
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonAreaResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return PokemonAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonAreaResp{}, err
	}

	pokemonAreaResp := PokemonAreaResp{}
	err = json.Unmarshal(data, &pokemonAreaResp)
	if err != nil {
		return PokemonAreaResp{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemonAreaResp, nil
}
