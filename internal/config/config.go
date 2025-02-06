package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct{
	Port uint32 `yaml:"port"`
	MaxDepth uint8 `yaml:"max_depth"`
	DisAgreeDomains []string `yaml:"disagree_domains"`
	LogsDir string `yaml:"logs_dir"`
}

func Get(configpath string) (*Config, error) {
	config := &Config{}
	file, err := os.ReadFile(configpath)
	if err != nil{
		return nil, err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil{
		return nil, err
	}	

	return config, nil
}