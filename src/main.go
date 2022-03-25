package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type Input struct {
	Location string
	Extended  bool
	Success  bool
	Error 	 string
}

func parseArgs(args []string) Input {
	usage := "Usage: weather <location> [extended]\nBy default output is short. Provides only the minimal weather information\n"
	if len(args) == 1 {
		return Input { Success: true, Location: args[0]}
			return Input { Success: false, Error: usage}
	} else if len(args) == 2 {
		if args[1] != "extended" {
			return Input { Success: false, Error: usage}
		} else {
			return Input { Success: true, Location: args[0], Extended: true}
		}
	} else {
		return Input { Success: false, Error: usage}
	}
}

func main() {
    	args := os.Args[1:]
    	input := parseArgs(args)
    	if input.Success == false {
    		fmt.Println(input.Error)
    		return
    	}

    	var config map[string]interface{}
    	data, err := os.ReadFile("config.txt",)
    	if err != nil {
    		fmt.Println("Failed to read config.txt. Please make sure you have it in the same folder as the executable")
    		return
    	}
    	json.Unmarshal(data, &config)
    	apiKey, ok := config["api_key"]
    	if !ok {
    		fmt.Println("API key for https://openweathermap.org has not been found in config.txt. Please add it.")
    		return
    	}

	query := queryLocation(input.Location, apiKey.(string))
	if !query.Success {
		fmt.Println(query.FailiureReason)
		return
	}

	weather := queryWeather(query.Result, Metric, apiKey.(string))
	if !weather.Success {
		fmt.Println("Failed to query weather data")
		return
	}
	fmt.Printf("%v°, %v\n", int(weather.Result.Main.Temp), weather.Result.Weather[0].Main)
	fmt.Printf("Feels like %v°\n", int(weather.Result.Main.Feels_like))
}
