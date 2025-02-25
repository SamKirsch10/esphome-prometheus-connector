package internal

import (
	"os"

	"gopkg.in/yaml.v3"
)

type metricYaml struct {
	Devices []struct {
		Name    string   `json:"name"`
		IP      string   `json:"ip"`
		Metrics []string `json:"metrics"`
	} `json:"devices"`
}

func loadYaml() metricYaml {
	f, err := os.ReadFile(YamlLocation)
	if err != nil {
		panic(err)
	}
	var data metricYaml
	if err := yaml.Unmarshal(f, &data); err != nil {
		panic(err)
	}
	return data
}
