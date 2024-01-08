package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Endpoints []Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	Name    string   `yaml:"name"`
	Method  string   `yaml:"method"`
	Path    string   `yaml:"path"`
	Handler string   `yaml:"handler"`
	Args    []string `yaml:"args"`
}

func Read(configFilePath string) (*Config, error) {
	yamlFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
