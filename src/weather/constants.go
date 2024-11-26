package weather

// Weather API Key
var weatherAPIKey string = GetWeatherAPIKey()

// Base OpenWeather URL
var baseOpenWeatherURL string = "https://api.openweathermap.org"

// OpenWeather URL to extract weather by city
var urlToExtractWeatherByCity string = baseOpenWeatherURL + "/data/2.5/weather?appid=" + weatherAPIKey + "&units=metric" + "&q="
