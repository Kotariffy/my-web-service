package config

import "sync"

type Config struct {
	DB DBconfig
}

type DBconfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

var (
	config *Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		config = &Config{}
	})

	return config
}
