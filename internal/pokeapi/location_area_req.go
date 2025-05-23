package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	//check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		result := LocationAreasResp{}
		err := json.Unmarshal(data, &result)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return result, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	result := LocationAreasResp{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return LocationAreasResp{}, err
	}
	//add data to cache
	c.cache.Add(fullURL, data)

	return result, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/"
	fullURL := baseURL + endpoint + locationAreaName

	//check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		result := LocationArea{}
		err := json.Unmarshal(data, &result)
		if err != nil {
			return LocationArea{}, err
		}
		return result, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	result := LocationArea{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return LocationArea{}, err
	}
	//add data to cache
	c.cache.Add(fullURL, data)

	return result, nil
}
