package server

import (
    "github.com/prometheus/client_golang/prometheus"
)


var (
    requestsCount = prometheus.NewCounter(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests.",
        },
    )
    requestDuration = prometheus.NewHistogram(
        prometheus.HistogramOpts{
            Name: "http_request_duration_seconds",
            Help: "Duration of HTTP requests.",
        },
    )
)

func init() {
    prometheus.MustRegister(requestsCount)
    prometheus.MustRegister(requestDuration)
}
