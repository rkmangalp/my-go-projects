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
```

## Obtain an API Key
 
 To use this application, you'll need an API key from [OpenWeatherMap](https://openweathermap.org/api). Follow these steps to obtain your API key:

1. Sign up for an account on [OpenWeatherMap](https://openweathermap.org/api).
2. Navigate to your API keys page after logging in.
3. Copy your API key. This key will be used to authenticate your requests to the OpenWeatherMap API.

## Update the Code with Your API Key

1. Open **main.go** in your preferred text editor.
2. Locate the line where **apiKey** is defined

```go
apiKey := "YOUR_API_KEY"
```

## Run the Application

To run the application, execute the following command in your terminal:

```sh
go run main.go
```

Upon running the application, you will see output similar to the following:

```yaml
Weather in London:
Temperature: 15.00°C
Weather Description: light rain
Humidity: 82%
Wind Speed: 5.10 m/s
```

## Customization

To fetch weather data for a different city, change the value of the **city** variable in **main.go**:

```go
city := "New York"
```
## Project Structure
 
 .
├── main.go
├── README.md

- main.go: The main Go file containing the logic to fetch and display weather data.
- README.md: This README file.

