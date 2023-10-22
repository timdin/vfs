package configs

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// enum values for storage mode
type Mode string

const (
	File Mode = "file"
	DB   Mode = "database"
)

// define config type
type Config struct {
	Mode Mode `yaml:"mode"`
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