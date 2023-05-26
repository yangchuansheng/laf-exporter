package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ApiResponse struct {
	From   string `json:"from"`
	Stats  Stats  `json:"stats"`
	QueryAt string `json:"query_at"`
}

type Stats struct {
	Applications       int `json:"applications"`
	RunningApplications int `json:"running_applications"`
	Users              int `json:"users"`
}

func FetchMetrics() {
    res, err := http.Get(config.URL + "?api_key=" + config.APIKey)
    if err != nil {
        log.Printf("Failed to fetch data: %v", err)
        return
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Printf("Failed to read response body: %v", err)
        return
    }

    var apiResponse ApiResponse
    err = json.Unmarshal(body, &apiResponse)
    if err != nil {
        log.Printf("Failed to parse JSON response: %v", err)
        return
    }

    // Update the metrics with labels
    applications.WithLabelValues(apiResponse.From).Set(float64(apiResponse.Stats.Applications))
    runningApplications.WithLabelValues(apiResponse.From).Set(float64(apiResponse.Stats.RunningApplications))
    users.WithLabelValues(apiResponse.From).Set(float64(apiResponse.Stats.Users))
}
