package metrics

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"log"
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

var (
	applications = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "applications",
			Help: "Total number of applications",
		},
		[]string{"source"},
	)
	runningApplications = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "running_applications",
			Help: "Total number of running applications",
		},
		[]string{"source"},
	)
	users = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "users",
			Help: "Total number of users",
		},
		[]string{"source"},
	)
)

func init() {
	prometheus.MustRegister(applications)
	prometheus.MustRegister(runningApplications)
	prometheus.MustRegister(users)
}

func UpdateMetrics(body []byte) {
    var apiResponse ApiResponse
    err := json.Unmarshal(body, &apiResponse)
    if err != nil {
        log.Printf("Failed to parse JSON response: %v", err)
        return
    }

    // Update the metrics with labels
    applications.WithLabelValues(apiResponse.From).Add(float64(apiResponse.Stats.Applications))
    runningApplications.WithLabelValues(apiResponse.From).Add(float64(apiResponse.Stats.RunningApplications))
    users.WithLabelValues(apiResponse.From).Add(float64(apiResponse.Stats.Users))
}
