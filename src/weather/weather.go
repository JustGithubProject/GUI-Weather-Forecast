package weather

func FetchWeatherForecastByCity(city string) string {
	temperature := doRequest(urlToExtractWeatherByCity + city)
	return temperature
}