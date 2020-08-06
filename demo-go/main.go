package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	cc := NewHttpRequestsTotalCount("contributor")
	cd := NewHttpRequestDurationHistogram("contributor")

	foundHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello from example application."))
	})
	notfoundHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	foundChain := promhttp.InstrumentHandlerDuration(
		cd.MustCurryWith(prometheus.Labels{"handler": "found"}),
		promhttp.InstrumentHandlerCounter(cc.MustCurryWith(prometheus.Labels{"handler": "found"}), foundHandler),
	)

	app := http.NewServeMux()
	metrics := http.NewServeMux()

	app.Handle("/", foundChain)
	app.Handle("/err", promhttp.InstrumentHandlerCounter(cc.MustCurryWith(prometheus.Labels{"handler": "notfound"}), notfoundHandler))

	metrics.Handle("/metrics", HandlerFuncForMetrics())
	go http.ListenAndServe("0.0.0.0:8001", app)
	log.Fatal(http.ListenAndServe("0.0.0.0:8101", metrics))
}
