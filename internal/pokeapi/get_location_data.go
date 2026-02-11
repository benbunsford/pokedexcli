package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetLocationData(areaName *string) (LocationData, error) {
	url := baseURL + "/location-area/" + *areaName

	if cached, ok := c.cache.Get(url); ok {
		data := LocationData{}
		err := json.Unmarshal(cached, &data)
		if err != nil {
			return LocationData{}, err
		}
		return data, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 ||
		resp.StatusCode < 200 {
		return LocationData{}, errors.New("No location data found, please try again!")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationData{}, err
	}

	c.cache.Add(url, body)

	data := LocationData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return LocationData{}, err
	}

	return data, nil
}
