package main

import (
	"context"
	"flag"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/samkirsch10/esphome-prometheus-connector/internal"
	log "github.com/sirupsen/logrus"
)

func main() {
	port := flag.String("port", "9000", "Port to listen on")
	lvl := flag.String("log-lvl", "INFO", "Log level")
	yamlLocation := flag.String("metrics-file", "/app/metrics.yaml", "The location of the metrics yaml file.")
	flag.Parse()

	loglvl := log.WarnLevel
	switch strings.ToUpper(*lvl) {
	case "INFO":
		loglvl = log.InfoLevel
	case "DEBUG":
		loglvl = log.DebugLevel
	case "WARN":
		loglvl = log.WarnLevel
	case "ERROR":
		loglvl = log.ErrorLevel
	default:
		panic("unknown log level. try `INFO`, `DEBUG`, `WARN`, or `ERROR`")
	}
	log.SetLevel(loglvl)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	internal.YamlLocation = *yamlLocation
	go internal.Run(ctx)

	router := mux.NewRouter()

	// Prometheus endpoint
	router.Path("/metrics").Handler(promhttp.Handler())

	log.Info("Serving requests on port " + *port)
	err := http.ListenAndServe(":"+*port, router)
	log.Fatal(err)
}
