package main

import (
	// "os"
	"log"
	// "JustGithubProject/GUI-Weather-Forecast/src/ui"
	"JustGithubProject/GUI-Weather-Forecast/src/weather"
)


func main() {
	// ui.RunWindow(len(os.Args), os.Args)
	log.Println(weather.FetchWeatherForecastByCity("Kair"))
}