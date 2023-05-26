package app

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	applications = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "applications",
			Help: "Total number of applications",
		},
		[]string{"source"},
	)
	runningApplications = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "running_applications",
			Help: "Total number of running applications",
		},
		[]string{"source"},
	)
	users = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
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
