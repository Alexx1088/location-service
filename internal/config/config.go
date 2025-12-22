package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"time"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Kafka  KafkaConfig  `yaml:"kafka"`
	Outbox OutboxConfig `yaml:"outbox"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
}

type OutboxConfig struct {
	BatchSize    int           `yaml:"batch_size"`
	Interval     time.Duration `yaml:"interval"`
	WorkersCount int           `yaml:"workers_count"`
}

func Load(path string) (*Config, error) {
	cfg := &Config{}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	env := os.Getenv("APP_ENV")
	var brokers string

	switch env {
	case "k8s":
		brokers = os.Getenv("KAFKA_BROKERS_K8S")
	case "minikube":
		brokers = os.Getenv("KAFKA_BROKERS_MINIKUBE")
	}

	if brokers != "" {
		cfg.Kafka.Brokers = strings.Split(brokers, ",")
	}

	return cfg, nil
}
