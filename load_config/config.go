package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	App struct {
		Port int `yaml:"port"`
	} `yaml:"app"`

	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"db"`

	Kafka struct {
		Brokers []string `yaml:"brokers"`
		GroupID string   `yaml:"group_id"`
		Topics  []string `yaml:"topics"`
	} `yaml:"kafka"`

	Email struct {
		APIKey string `yaml:"api_key"`
		From   string `yaml:"from"`
	} `yaml:"email"`
}

func LoadConfig(path string) *AppConfig {
	data, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	var cfg AppConfig

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("error parsing config: %v", err)
	}

	return &cfg
}
