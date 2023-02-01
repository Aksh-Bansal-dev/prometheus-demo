package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	reqCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "home_req_total",
		Help: "Total number of requests to homepage",
	})
)

func main() {
	http.HandleFunc("/ip/", func(w http.ResponseWriter, r *http.Request) {
        reqCount.Inc()
		w.Write([]byte(`{"status": "ok"}`))
	})

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", http.FileServer(http.Dir("./static")))
    log.Println("Server started on port: 2001")
	http.ListenAndServe(":2001", nil)
}
