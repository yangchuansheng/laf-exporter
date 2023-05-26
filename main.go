package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/prometheus/client_golang/prometheus/promhttp"
    "laf-exporter/app"
)

func main() {
    app.LoadConfig()  // Load the config

    // Set up a ticker that triggers every 30 seconds
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()

    go func() {
        for {
            app.FetchMetrics()  // Call fetchMetrics to populate the metrics
            <-ticker.C // Wait for the next tick
        }
    }()

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
