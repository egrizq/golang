package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	var key string = "8fba757e84fe4783e2bb8ab93aa117c3"
	cities := []string{"Bekasi", "Jakarta", "Bali"}

	for _, city := range cities {
		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", city, key)

		response, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close()

		if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
			log.Fatal(err)
		}

		temprature := data.Main.Temp - 273.15
		fmt.Println("temprature:", math.Round(temprature), "°C")
	}

	fmt.Println("total duration:", time.Since(start))
}
