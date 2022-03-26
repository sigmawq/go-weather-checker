package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type Input struct {
	Location string
	Units  	 Units
	Extended  bool
	Error 	 string

	Success  bool
}

func parseArgs(args []string) Input {
	usage := "Usage: weather <location> [metric | imperial] [extended]\nBy default output is short. Provides only the minimal weather information\n Default unit is metric"
	var input Input

	for _, arg := range args {
		switch arg {
		case "metric":
			input.Units = Metric
		case "imperial":
			input.Units = Imperial
		case "extended":
			input.Extended = true
		default:
			input.Location = arg
		}
	}

	if input.Location == "" {
		input.Success = false
		input.Error = usage
	} else {
		input.Success = true
	}

	return input
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

	if input.Extended {
		if input.Units == Imperial {
			fmt.Printf("Wind %v miles/h %v°\n", int(weather.Result.Wind.Speed), int(weather.Result.Wind.Deg))	
		} else {
			fmt.Printf("Wind %v meters/h %v°\n", int(weather.Result.Wind.Speed), int(weather.Result.Wind.Deg))	
		}
		fmt.Printf("Atmospheric pressure %vhPa\n", int(weather.Result.Main.Pressure))
		fmt.Printf("Humidity %v%%\n", int(weather.Result.Main.Humidity))
		fmt.Printf("Clouds %v%%\n", int(weather.Result.Clouds.All))
	}
}
