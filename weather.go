package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Структура для ответа OpenWeatherMap
type WeatherResponse struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
}

// GetWeather запрашивает погоду для указанного города
func GetWeather(city string) (temp float64, humidity int, err error) {
	apiKey := os.Getenv("WEATHER_API_KEY") // ключ из переменной окружения
	if apiKey == "" {
		return 0, 0, fmt.Errorf("WEATHER_API_KEY not set")
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, 0, fmt.Errorf("API error: %s", resp.Status)
	}

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return 0, 0, fmt.Errorf("JSON decode failed: %w", err)
	}

	return weather.Main.Temp, weather.Main.Humidity, nil
}
