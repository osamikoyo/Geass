package config

type Config struct{
	Port uint32 `yaml:"port"`
	MaxDepth uint8 `yaml:"max_depth"`
	DisAgreeDomains []string `yaml:"disagree_domains"`
}

func New()