package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	//check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		result := Pokemon{}
		err := json.Unmarshal(data, &result)
		if err != nil {
			return Pokemon{}, err
		}
		return result, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	result := Pokemon{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return Pokemon{}, err
	}
	//add data to cache
	c.cache.Add(fullURL, data)

	return result, nil
}
