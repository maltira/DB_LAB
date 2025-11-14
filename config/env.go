package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUser string
	DbPass string
	DbHost string
	DbPort string
	DbName string
}

var Env *Config

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Не удалось получить доступ к .env файлу")
	}
	Env = &Config{
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASS"),
		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbName: os.Getenv("DB_NAME"),
	}
}
