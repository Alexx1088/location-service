package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	Outbox OutboxConfig `yaml:"outbox"`
}

type OutboxConfig struct {
	BatchSize int           `yaml:"batch_size"`
	Interval  time.Duration `yaml:"interval"`
}

func Load(path string) (*Config, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(raw, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
