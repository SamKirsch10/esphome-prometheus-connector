package internal

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	temperatureSensor = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "esp32_temperature_sensor_celsius",
		},
		[]string{"name"},
	)
)

func init() {
	prometheus.Register(temperatureSensor)
}
