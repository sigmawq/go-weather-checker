# Weather check
A simple tool for checking weather in any given location. Calls to https://openweathermap.org/

## Build and use
1. `go build -o weather.exe src/`
2. Create `config.txt` in the same folder as the executable.
3. Get an API key from https://openweathermap.org/
4. Specify it as the value of `api_key` field. Config itself is in JSON format. 

`config.txt` that is located in this repo is config template. You can use it.

### Usage
`weather <location> [metric | imperial] [extended]`
By default the app provides only a short and basic weather description.
Exetended version provides more info.
Default unit of measure is metric.