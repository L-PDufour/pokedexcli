package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(name string) (PokemonResp, error) {
	endpoint := "/pokemon/"

	fullURL := baseURL + endpoint + name

	if cachedRespBytes, ok := c.cache.Get(fullURL); ok {
		var cachedResp PokemonResp
		if err := json.Unmarshal(cachedRespBytes, &cachedResp); err != nil {
			return PokemonResp{}, fmt.Errorf("failed to unmarshal cached response: %v", err)
		}
		return cachedResp, nil
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return PokemonResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResp{}, err
	}

	pokemonResp := PokemonResp{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return PokemonResp{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemonResp, nil
}
