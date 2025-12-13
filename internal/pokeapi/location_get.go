package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(location_name string) (Location, error) {
	url := baseURL + "/location-area/" + location_name

	if val, ok := c.cache.Get(url); ok {
		pokemonList := Location{}
		err := json.Unmarshal(val, &pokemonList)
		if err != nil {
			return pokemonList, err
		}
		return pokemonList, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, nil
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	pokemonList := Location{}
	err = json.Unmarshal(dat, &pokemonList)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)

	return pokemonList, nil
}
