package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

var YamlLocation string

func Run(ctx context.Context) {
	d := loadYaml()

	t := time.NewTicker(10 * time.Second)
	gather(d)
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			gather(d)
		}
	}
}

func gather(metrics metricYaml) {
	cl := http.Client{}

	type metricData struct {
		ID    string  `json:"id"`
		Value float64 `json:"value"`
		State string  `json:"state"`
	}

	for _, device := range metrics.Devices {
		for _, m := range device.Metrics {
			log.Debugf("querying for metric %s from %s", device.IP, m)
			resp, err := cl.Get(fmt.Sprintf("http://%s/%s", device.IP, m))
			if err != nil {
				log.Errorf("failed to get metrics from %s", resp.Request.URL)
				continue
			}
			defer resp.Body.Close()
			var d metricData
			if err = json.NewDecoder(resp.Body).Decode(&d); err != nil {
				log.Errorf("failed to read response: %v", err)
				continue
			}
			log.Debugf("got response: '%s: %f'", d.ID, d.Value)
			name := strings.Split(d.ID, strings.Split(m, "/")[0]+"-")[1]
			temperatureSensor.WithLabelValues(name).Set(d.Value)
		}
	}
}
