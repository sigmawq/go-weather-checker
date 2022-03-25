package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Coord struct {
	Lon float64
	Lat float64
}

type Weather struct {
	Id          int
	Main        string
	Description string
	Icon        string
}

type MainData struct {
	Temp       float64
	Feels_like float64
	Temp_min   float64
	Temp_max   float64
	Pressure   float64
	Humidity   float64
}

type Wind struct {
	Speed float64
	Deg   float64
	Gust  float64
}

type Clouds struct {
	All float64
}

type Rain struct {
	One_hr   float64 `json:"1h"`
	Three_hr float64 `json:"3h"`
}

type Snow struct {
	One_hr   float64 `json:"1h"`
	Three_hr float64 `json:"3h"`
}

type Sys struct {
	Id      float64
	Message float64
	Country string
	Sunrise int
	Sunset  int
}

type WeatherOutput struct {
	Coord      Coord
	Weather    []Weather `json:"weather"`
	Base       string
	Main       MainData
	Visibility float64
	Wind       Wind
	Clouds     Clouds
	Rain       Rain
	Dt         int
	Sys        Sys
	Timezone   float64
	Id         int
	Name       string
	Cod        int
}

type WeatherQueryResult struct {
	Result         WeatherOutput
	Success        bool
	FailiureReason string
}

type Units int

const (
	Metric Units = iota
	Imperial
)

func queryWeather(location Location, units Units, apiKey string) WeatherQueryResult {
	var url strings.Builder

	url.WriteString("https://api.openweathermap.org/data/2.5/weather?lat=")
	url.WriteString(fmt.Sprintf("%f", location.Lat))
	url.WriteString("&lon=")
	url.WriteString(fmt.Sprintf("%f", location.Lon))
	if units == Metric {
		url.WriteString("&units=metric")
	} else {
		url.WriteString("&units=imperial")
	}
	url.WriteString("&appid=")
	url.WriteString(apiKey)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url.String(), nil)

	if err != nil {
		fmt.Println(err)
		return WeatherQueryResult{Success: false, FailiureReason: "HTTP request failed, check your internet connection"}
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return WeatherQueryResult{Success: false, FailiureReason: "HTTP request failed, check your internet connection"}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return WeatherQueryResult{Success: false, FailiureReason: "Read failed"}
	}

	var responseRaw map[string]interface{}
	json.Unmarshal(body, &responseRaw)
	errorMessage, ok := checkResponse(responseRaw)
	if !ok {
		return WeatherQueryResult{Success: false, FailiureReason: errorMessage}
	}

	var weatherOutput WeatherOutput
	json.Unmarshal([]byte(body), &weatherOutput)

	return WeatherQueryResult{Result: weatherOutput, Success: true}
}
