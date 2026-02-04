package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetMapData(reqUrl *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if reqUrl != nil {
		url = *reqUrl
	}

	if cached, ok := c.cache.Get(url); ok {
		data := RespShallowLocations{}
		err := json.Unmarshal(cached, &data)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return data, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	c.cache.Add(url, body)

	data := RespShallowLocations{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return data, nil
}
