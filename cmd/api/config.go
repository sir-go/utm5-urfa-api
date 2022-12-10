package main

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type (
	cfgBilling struct {
		Addr string `yaml:"addr"`
		Cert string `yaml:"cert"`
	}
	billingsMap map[string]cfgBilling

	Config struct {
		Billings billingsMap `yaml:"billings"`
	}
)

// readConfig parses Config struct from yaml data
func readConfig(confData []byte) (*Config, error) {
	conf := new(Config)
	if err := yaml.Unmarshal(confData, conf); err != nil {
		return nil, err
	}
	return conf, nil
}

// LoadConfig reads yaml file and parses it to Config struct
func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	return readConfig(data)
}
