package config

import (
	"log"
	"os"

	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisURL    string
	SqlUserName string
	SqlPassword string
	SqlHost     string
	SqlDatabase string
	SqlPort     string
}

func defaultConfig() *Config {
	return &Config{
		RedisURL:    os.Getenv("REDIS_URL"),
		SqlUserName: os.Getenv("SQL_USERNAME"),
		SqlPassword: os.Getenv("SQL_PASSWORD"),
		SqlHost:     os.Getenv("SQL_HOST"),
		SqlDatabase: os.Getenv("SQL_DATABASE"),
		SqlPort:     os.Getenv("SQL_PORT"),
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
