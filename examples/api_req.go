package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	versalog "github.com/VersaLog/VersaLog.go/VersaLog"
)

type WeatherResponse struct {
	Name    string `json:"name"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
		Pressure int     `json:"pressure"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

func main() {
	logger := versalog.NewVersaLog("detailed", false, true, "Request", false, false, false, []string{}, false, false)

	api := "http://api.openweathermap.org/data/2.5/weather"
	params := "?q=location name&appid=api key&units=metric&lang=ja"
	url := api + params

	resp, err := http.Get(url)
	if err != nil {
		logger.Error(fmt.Sprintf("failed: %v", err))
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		var data WeatherResponse
		if err := json.Unmarshal(body, &data); err != nil {
			logger.Error(fmt.Sprintf("JSON parse error: %v", err))
			os.Exit(1)
		}
		logger.Info("success")
		msg := fmt.Sprintf("< %sの天気予報 >\n\n> 天気\n・%s\n\n> 気温\n・%.1f°C\n\n> 湿度\n・%d%%\n\n> 気圧\n・%d hPa\n\n> 風速\n・%.1f m/s",
			data.Name,
			data.Weather[0].Description,
			data.Main.Temp,
			data.Main.Humidity,
			data.Main.Pressure,
			data.Wind.Speed,
		)
		fmt.Println(msg)
	} else {
		logger.Error(fmt.Sprintf("failed: status code %d", resp.StatusCode))
	}
}
