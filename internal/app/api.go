package app

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

const (
	cityEndpoint    string = "https://geocoding-api.open-meteo.com/v1/search?name=%s&language=ru"
	weatherEndpoint string = "https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true"
)

var (
	ErrCityNotFound    = errors.New("город не найден")
	ErrWeatherNotFound = errors.New("не удалось определить погоду")
)

type GeoResultElement struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeoResult struct {
	Results []GeoResultElement `json:"results"`
}

type CurrentWeatherResult struct {
	Name          string
	Temperature   float64 `json:"temperature"`
	WindSpeed     float64 `json:"windspeed"`
	WindDirection float64 `json:"winddirection"`
}

type WeatherResult struct {
	CurrentWeather CurrentWeatherResult `json:"current_weather"`
	Reason         string               `json:"reason"`
}

func GetWeatherByCoordinates(latitude float64, longitude float64) (CurrentWeatherResult, error) {
	var result WeatherResult
	client := resty.New()
	_, err := client.R().SetResult(&result).Get(fmt.Sprintf(weatherEndpoint, latitude, longitude))
	if err != nil {
		return CurrentWeatherResult{}, err
	}

	if result.Reason != "" {
		return CurrentWeatherResult{}, ErrWeatherNotFound
	}

	return result.CurrentWeather, nil
}

func GetGeographicalCoordinates(city string) (GeoResultElement, error) {
	var result GeoResult
	client := resty.New()
	_, err := client.R().SetResult(&result).Get(fmt.Sprintf(cityEndpoint, city))
	if err != nil {
		return GeoResultElement{}, err
	}

	if result.Results == nil {
		return GeoResultElement{}, ErrCityNotFound
	}

	return result.Results[0], nil
}
