package configs

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// enum values for mode
type Mode string

const (
	Prod Mode = "prod"
	Dev  Mode = "dev"
)

type DBConfig struct {
	Conn string `yaml:"conn"`
}

// define config type
type Config struct {
	Mode     Mode     `yaml:"mode"`
	DBconfig DBConfig `yaml:"database"`
}

var AppConfig *Config

func LoadConfig() {
	AppConfig = &Config{}
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("current mode: ", AppConfig.Mode)
}
