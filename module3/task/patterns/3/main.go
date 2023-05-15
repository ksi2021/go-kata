package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func UnmarshalTest(data []byte) (Test, error) {
	var r Test
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Test) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Test struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Current struct {
	TempC     float64   `json:"temp_c"`
	Condition Condition `json:"condition"`
	WindKph   float64   `json:"wind_kph"`
	Humidity  float64   `json:"humidity"`
	Uv        float64   `json:"uv"`
}

type Condition struct {
}

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}

// WeatherAPI is the interface that defines the methods for accessing weather information
type WeatherAPI interface {
	GetTemperature(location string) int
	GetHumidity(location string) int
	GetWindSpeed(location string) int
}

// OpenWeatherAPI is the implementation of the weather API
type OpenWeatherAPI struct {
	apiKey string
}

func (o *OpenWeatherAPI) GetTemperature(location string) int {
	// Make a request to the open weather API to retrieve temperature information
	// and return the result
	// ...
	str := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", o.apiKey, url.QueryEscape(location))
	// fmt.Println(str)
	resp, err := http.Get(str)
	if err != nil {
		panic(err)
	}
	data, _ := io.ReadAll(resp.Body)
	t, er := UnmarshalTest(data)

	if er != nil {
		panic(er)
	}

	return int(t.Current.TempC)
}

func (o *OpenWeatherAPI) GetHumidity(location string) int {
	// Make a request to the open weather API to retrieve humidity information
	// and return the result
	// ...

	str := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", o.apiKey, url.QueryEscape(location))

	resp, err := http.Get(str)
	if err != nil {
		panic(err)
	}
	data, _ := io.ReadAll(resp.Body)
	t, er := UnmarshalTest(data)
	if er != nil {
		panic(er)
	}

	return int(t.Current.Humidity)
}

func (o *OpenWeatherAPI) GetWindSpeed(location string) int {
	// Make a request to the open weather API to retrieve wind speed information
	// and return the result
	// ...
	str := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", o.apiKey, url.QueryEscape(location))

	resp, err := http.Get(str)
	if err != nil {
		panic(err)
	}
	data, _ := io.ReadAll(resp.Body)
	t, er := UnmarshalTest(data)
	if er != nil {
		panic(er)
	}

	return int(t.Current.WindKph)

}

// WeatherFacade is the facade that provides a simplified interface to the weather API
type WeatherFacade struct {
	weatherAPI WeatherAPI
}

func (w *WeatherFacade) GetWeatherInfo(location string) (int, int, int) {
	temperature := w.weatherAPI.GetTemperature(location)
	humidity := w.weatherAPI.GetHumidity(location)
	windSpeed := w.weatherAPI.GetWindSpeed(location)

	return temperature, humidity, windSpeed
}

func NewWeatherFacade(apiKey string) WeatherFacade {
	return WeatherFacade{
		weatherAPI: &OpenWeatherAPI{apiKey: apiKey},
	}
}

func main() {
	weatherFacade := NewWeatherFacade("f1bb0056ceba49278b5213744231505")
	cities := []string{"Москва", "Санкт-Петербуг", "Казань", "Якутск"}

	for _, city := range cities {
		temperature, humidity, windSpeed := weatherFacade.GetWeatherInfo(city)
		fmt.Printf("Temperature in "+city+": %d\n", temperature)
		fmt.Printf("Humidity in "+city+": %d\n", humidity)
		fmt.Printf("Wind speed in "+city+": %d\n\n", windSpeed)
	}

	// 9e4e114128cb5762e995e0abd0216124

	// resp, err := http.Get(`http://api.openweathermap.org/data/2.5/weather?q=London,uk&units=metric&APPID=716e25cdfd1e6b22bc2d21d1f916341c`)
	// if err != nil {
	// 	panic(err)
	// }
	// data, _ := io.ReadAll(resp.Body)
	// // fmt.Println(resp.StatusCode, data, string(data))
	// t, er := UnmarshalTest(data)
	// fmt.Printf("%v \t %+v ", er, t.Main)

	// data := []byte(`{"coord":{"lon":-0.1257,"lat":51.5085},"weather":[{"id":801,"main":"Clouds","description":"few clouds","icon":"02n"}],"base":"stations","main":{"temp":10.4,"feels_like":9.35,"temp_min":8.36,"temp_max":11.86,"pressure":1022,"humidity":71},"visibility":10000,"wind":{"speed":2.06,"deg":60},"clouds":{"all":18},"dt":1684184763,"sys":{"type":2,"id":2075535,"country":"GB","sunrise":1684123758,"sunset":1684179874},"timezone":3600,"id":2643743,"name":"London","cod":200}`)
	// // t, e := UnmarshalTest(data)
	// _ = data

	// dd, err := UnmarshalTest(data)

	// fmt.Printf("%+v \t \n \n %v", dd.Main.Temp, err)
}
