# Go Weather Application

This Go Weather Application fetches real-time weather data from the OpenWeatherMap API and displays it in the console. This project demonstrates how to interact with external APIs, parse JSON responses, and handle HTTP requests in Go.

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or later)
- An API key from [OpenWeatherMap](https://openweathermap.org/api)

## Getting Started

### Clone the Repository

```sh
git clone https://github.com/your-username/go-weather-application.git
cd go-weather-application

## Obtain an API Key

To use this application, you'll need an API key from [OpenWeatherMap](https://openweathermap.org/). Follow these steps to obtain your API key:

1. **Sign up** for an account on [OpenWeatherMap](https://openweathermap.org/).
2. Navigate to your **API keys** page after logging in.
3. Copy your **API key**. This key will be used to authenticate your requests to the OpenWeatherMap API.

## Update the code with your API key

Open **`main.go`** and replace **`YOUR_API_KEY`** with your actual API key:

```sh
apiKey := "Your_API_KEY" 