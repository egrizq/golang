package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	resultCh := make(chan float64)
	var wg sync.WaitGroup

	cityList := []string{"Jakarta", "Bekasi", "Bali"}

	for _, cities := range cityList {
		wg.Add(1)
		go timer(cities, resultCh, &wg)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// printing temprature
	for temp := range resultCh {
		fmt.Println("temprature:", math.Round(temp), "°C")
	}

	duration := time.Since(start)
	fmt.Println("total duration:", duration)
}

func timer(cities string, resultCh chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()

	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	var api string = "https://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s"
	var key string = "8fba757e84fe4783e2bb8ab93aa117c3"
	url := fmt.Sprintf(api, cities, key)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}

	temprature := data.Main.Temp - 273.15
	resultCh <- temprature
}
