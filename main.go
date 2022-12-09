package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	apiUrl, exists := os.LookupEnv("API_URL")
	if !exists {
		apiUrl = "http://localhost:26657"
	}
	listenIp, exists := os.LookupEnv("LISTEN_IP")
	if !exists {
		listenIp = "0.0.0.0:9000"
	}

	var (
		status  Status
		netinfo Netinfo
	)

	go func() {
		for {
			status = queryStatus(apiUrl)
			netinfo = queryNetinfo(apiUrl)

			time.Sleep(2 * time.Second)
		}
	}()

	if err := prometheus.Register(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace:   "gaiad",
			Name:        "latest_block_height",
			Help:        "Latest Block Height",
			ConstLabels: prometheus.Labels{"destination": "primary"},
		},
		func() float64 { return float64(status.latestBlockHeight) },
	)); err != nil {
		log.Fatal(err)
	}

	if err := prometheus.Register(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace:   "gaiad",
			Name:        "latest_block_time",
			Help:        "Latest Block Time",
			ConstLabels: prometheus.Labels{"destination": "primary"},
		},
		func() float64 { return float64(status.latestBlockTime) },
	)); err != nil {
		log.Fatal(err)
	}

	if err := prometheus.Register(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace:   "gaiad",
			Name:        "block_desync_time",
			Help:        "Block Desync Time",
			ConstLabels: prometheus.Labels{"destination": "primary"},
		},
		func() float64 { return float64(status.blockDesyncTime) },
	)); err != nil {
		log.Fatal(err)
	}

	if err := prometheus.Register(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Namespace:   "gaiad",
			Name:        "peers_count",
			Help:        "Peers Count",
			ConstLabels: prometheus.Labels{"destination": "primary"},
		},
		func() float64 { return float64(netinfo.peersCount) },
	)); err != nil {
		log.Fatal(err)
	}

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(listenIp, nil)
}
