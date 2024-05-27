package config

import (
	"log"
	"os"

	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisURL string
}

func defaultConfig() *Config {
	return &Config{
		RedisURL: os.Getenv("REDIS_URL"),
	}
}

var (
	_config       *Config
	configFactory = defaultConfig
)

func GetConfig() *Config {
	if _config == nil {
		_config = configFactory()
		return _config
	}
	return _config
}

func Init() {
	dir, _ := os.Getwd()
	basePath := filepath.Join(dir, ".env")

	if err := godotenv.Load(basePath); err != nil {
		log.Fatal("Error Loading .env. file")
	}
}
