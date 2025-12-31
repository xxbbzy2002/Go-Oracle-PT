package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Oracle OracleConfig `yaml:"oracle"`
	Log    LogConfig    `yaml:"log"`
}

type OracleConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Service  string `yaml:"service"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type LogConfig struct {
	Dir string `yaml:"dir"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
