package weather

import (
	"fmt"
	"log"
	"os"
	"io"
	"net/http"
	"encoding/json"
	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func GetWeatherAPIKey() string {
	/*
		Extracting API Key from .env 
	*/
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to find .env file")
	}

	return os.Getenv("WEATHER_API_KEY")
}


func doRequest(URL string) string {
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal("Failed to do request to" + URL)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Failed to read response")
	}

	// Parsing JSON response
	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatal("Failed to parse JSON response")
	}

	temperature := weather.Main.Temp

	return fmt.Sprintf("Temperature: %.2f°C", temperature)

}