package main

import (
    "fmt"
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"laf-exporter/internal/api"
    "laf-exporter/internal/config"
    "laf-exporter/internal/metrics"
)

func main() {
    config.LoadConfig()  // Load the config

    // Set up a ticker that triggers every 30 seconds
    ticker := api.StartTicker(30)
    defer ticker.Stop()

    // Fetch metrics every 30 seconds
    body, err := api.FetchMetrics()  // Fetch metrics from the API
    if err != nil {
        log.Printf("Failed to fetch metrics: %v", err)
        return
    }

    metrics.UpdateMetrics(body)  // Update the metrics with the fetched data

    // Set up endpoints
    http.Handle("/metrics", promhttp.Handler())

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        // When this endpoint is hit, respond with a simple message indicating that the server is up
        fmt.Fprintf(w, "Server is up and running!")
    })

    log.Println("Server is starting...")

    go func() {
        if err := http.ListenAndServe(":8080", nil); err != nil {
            log.Fatalf("Failed to start server: %v", err)
        }
    }()

    log.Println("Server started successfully, listening on port 8080...")
    select {} // Keep the main function from returning, which would end the program
}
