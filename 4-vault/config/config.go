package config

import (
	"fmt"
	"os"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	envKey := os.Getenv("KEY")
	if envKey == "" {
		fmt.Println("KEY environment variable not set")
		return nil
	}
	return &Config{
		Key: envKey,
	}
}
