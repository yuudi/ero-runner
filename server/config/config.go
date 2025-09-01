package config

import (
	"log"
	"os"
	"sync"
)

type configType struct {
	Listen string
}

var config *configType

func loadConfigFromEnv() {
	config = &configType{
		Listen: getEnv("ERO_LISTEN_ADDRESS", ":8080"),
	}
}

var once sync.Once

func GetConfig() *configType {
	once.Do(loadConfigFromEnv)
	return config
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Required environment variable %s is not set", key)
	}
	return value
}
