package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetMapData(reqUrl *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if reqUrl != nil {
		url = *reqUrl
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

	data := RespShallowLocations{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return data, nil
}
