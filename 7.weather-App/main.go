package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Define a struct to represent the JSON response from the API
type WeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`

	// Slice of structs to hold weather information
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`

	Base string `json:"base"`

	// Nested struct to hold main weather details
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`

	// Nested struct to hold wind information
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`

	Clouds struct { // Nested struct to hold cloud information
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int      `json:"dt"`
	Sys struct { // Nested struct to hold system information
		Type    int    `json:"type"`
		Id      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func main() {

	apiKey := "Your_API_KEY" // Replace 'YOUR_API_KEY' with your OpenWeatherMap API key

	// promt user to enter the city
	var city string
	fmt.Println("Enter the city name")
	_, err := fmt.Scan(&city)
	if err != nil {
		fmt.Printf("Error reading input: %s\n", err)
		return
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	// Make a GET request to the API
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data: %s\n", err)
		return
	}

	defer response.Body.Close() // Ensure the response body is closed after function exits

	// Read the response

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading the response body: %s\n", err)
		return
	}

	// Parse JSON response into WeatherResponse struct
	var weatherResp WeatherResponse
	err = json.Unmarshal(body, &weatherResp)
	if err != nil {
		fmt.Printf("Error while parsing the body: %s\n", err)
		return
	}

	// Display weather information
	fmt.Printf("Weather in %s: \n", weatherResp.Name)
	fmt.Printf("Temperature: %.2f\n", weatherResp.Main.Temp)
	fmt.Printf("Description: %s\n", weatherResp.Weather[0].Description)
	fmt.Printf("Humidity: %d%%\n", weatherResp.Main.Humidity)
	fmt.Printf("Wind Speed: %.2f\n", weatherResp.Wind.Speed)
}
