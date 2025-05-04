package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBHost                 string
	DBPort                 string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", "8080"),
		DBUser:                 getEnv("DB_USER", "postgres"),
		DBPassword:             getEnv("DB_PASSWORD", "postgres"),
		DBHost:                 getEnv("DB_HOST", "127.0.0.1"),
		DBPort:                 getEnv("DB_PORT", "5432"),
		DBName:                 getEnv("DB_NAME", "GIN_DATABASE"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret:              getEnv("JWTSecret", "not-secret"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
