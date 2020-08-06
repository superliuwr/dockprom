package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewCounter(subsystem string, name string, help string, labelNames []string) *prometheus.CounterVec {
	// TODO Check parameters are valid

	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: subsystem,
		Name:      name,
		Help:      help,
	}, labelNames)

	// TODO Don't panic when fail
	prometheus.MustRegister(counter)

	return counter
}

// func NewGauge(subsystem string, name string, help string, labelNames []string) *prometheus.GaugeVec {
// 	// TODO Check parameters are valid
//
// 	gauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 		Subsystem: subsystem,
// 		Name: name,
// 		Help: help,
// 	}, labelNames)
//
// 	// TODO Don't panic when fail
// 	registry.MustRegister(gauge)
//
// 	return gauge
// }

func NewHistogram(subsystem string, name string, help string, buckets []float64, labelNames []string) *prometheus.HistogramVec {
	// TODO Check parameters are valid

	histogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: subsystem,
		Name:      name,
		Help:      help,
		Buckets:   buckets,
	}, labelNames)

	// TODO Don't panic when fail
	prometheus.MustRegister(histogram)

	return histogram
}

func NewHttpRequestsTotalCount(component string) *prometheus.CounterVec {
	return NewCounter(component, "http_requests_total", "Count of all HTTP requests", []string{"code", "method", "handler"})
}

func NewHttpRequestDurationHistogram(component string) *prometheus.HistogramVec {
	return NewHistogram(component, "http_requests_duration_seconds", "Duration of all HTTP requests in seconds", prometheus.DefBuckets, []string{"code", "method", "handler"})
}

func HandlerFuncForMetrics() http.Handler {
	return promhttp.Handler()
}
