package config_reader

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type YAMLConfigReader struct {
}

func (ycr *YAMLConfigReader) GetConfig() (*Config, error) {
	filename := filepath.Join("config", "cfg.yaml")

	if _, err := os.Stat(filename); err != nil {
		log.Panicf("ERROR: %v\n", err.Error())
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicf("ERROR: %v\n", err.Error())
	}

	cfg := Config{}
	if err := yaml.Unmarshal(content, &cfg); err != nil {
		log.Panicf("ERROR: %v\n", err)
	}

	return &cfg, nil
}

func NewYAMLConfigReader() (ConfigReader, error) {
	return &YAMLConfigReader{}, nil
}
