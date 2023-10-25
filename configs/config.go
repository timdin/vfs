package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

// enum values for mode
type Mode string

const (
	Prod Mode = "prod"
	Dev  Mode = "dev"
)

type StorageType string

const (
	RemoteStorage StorageType = "remote"
	LocalStorage  StorageType = "local"
)

type MySqlConfig struct {
	Conn string `yaml:"conn"`
}

type SqliteConfig struct {
	Path string `yaml:"path"`
}

// define config type
type Config struct {
	DBmode         Mode         `yaml:"db_mode"`
	DBtype         StorageType  `yaml:"db_type"`
	RemoteDBconfig MySqlConfig  `yaml:"database"`
	LocalDBconfig  SqliteConfig `yaml:"local_database"`
}

func LoadConfig() *Config {
	c := &Config{}
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		panic(err)
	}
	return c
}
