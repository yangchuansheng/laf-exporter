package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"laf-exporter/internal/config"
)


func StartTicker(durationInSeconds int) *time.Ticker {
	// Set up a ticker that triggers every specified seconds
	ticker := time.NewTicker(time.Duration(durationInSeconds) * time.Second)
	go func() {
		for {
			FetchMetrics()  // Call FetchMetrics to populate the metrics
			<-ticker.C // Wait for the next tick
		}
	}()
	return ticker
}

func FetchMetrics() ([]byte, error) {
    config := config.GetConfig()

    res, err := http.Get(config.URL + "?api_key=" + config.APIKey)
    if err != nil {
		log.Printf("Failed to fetch data: %v", err)
        return nil, err
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
		log.Printf("Failed to read response body: %v", err)
        return nil, err
    }

    log.Printf("Received response: %s", string(body))

    // Return the response body to be processed elsewhere
    return body, nil
}
