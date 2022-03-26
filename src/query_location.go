package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "strings"
)

type Location struct {
	Message string
	Lon float64
	Lat float64
}

type LocationQueryResult struct {
	Result         Location
	Success        bool
	FailiureReason string
}

func queryLocation(city string, apiKey string) LocationQueryResult {
	url := "http://api.openweathermap.org/geo/1.0/direct?q="
	url += city
	url += "&appid="
	url += apiKey
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return LocationQueryResult{Success: false, FailiureReason: "HTTP request failed, check your internet connection"}
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return LocationQueryResult{Success: false, FailiureReason: "HTTP request failed, check your internet connection"}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return LocationQueryResult{Success: false, FailiureReason: "Read failed"}
	}

	var responseRaw map[string]interface{}
	json.Unmarshal(body, &responseRaw)
	errorMessage, ok := checkResponse(responseRaw)
	if !ok {
		return LocationQueryResult{Success: false, FailiureReason: errorMessage}
	}

	var locations []Location
	json.Unmarshal(body, &locations)

	if len(locations) == 0 {
		return LocationQueryResult{Success: false, FailiureReason: "Location not found"}
	}
	location := locations[0]

	return LocationQueryResult{Result: location, Success: true}
}
