package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetPokemonData(pokemonName *string) (PokemonData, error) {
	url := baseURL + "/pokemon/" + *pokemonName

	if cached, ok := c.cache.Get(url); ok {
		data := PokemonData{}
		err := json.Unmarshal(cached, &data)
		if err != nil {
			return PokemonData{}, err
		}
		return data, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 ||
		resp.StatusCode < 200 {
		return PokemonData{}, errors.New("No pokemon data found, please try again!")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonData{}, err
	}

	c.cache.Add(url, body)

	data := PokemonData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return PokemonData{}, err
	}

	return data, nil
}
