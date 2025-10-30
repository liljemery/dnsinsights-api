package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName   string
	AppEnv    string
	AppPort   int
	DBHost    string
	DBPort    int
	DBUser    string
	DBPass    string
	DBName    string
	DBSSLMode string
	LogLevel  string
	RateLimit int
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	port := getInt("APP_PORT", 8080)
	dbPort := getInt("DB_PORT", 5432)
	rate := getInt("RATE_LIMIT", 100)

	cfg := &Config{
		AppName:   getStr("APP_NAME", "DNS Insight API"),
		AppEnv:    getStr("APP_ENV", "development"),
		AppPort:   port,
		DBHost:    getStr("DB_HOST", "localhost"),
		DBPort:    dbPort,
		DBUser:    getStr("DB_USER", "postgres"),
		DBPass:    getStr("DB_PASSWORD", "postgres"),
		DBName:    getStr("DB_NAME", "dnsinsight"),
		DBSSLMode: getStr("DB_SSLMODE", "disable"),
		LogLevel:  getStr("LOG_LEVEL", "info"),
		RateLimit: rate,
	}
	return cfg, nil
}

func (c *Config) PostgresDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC",
		c.DBHost, c.DBUser, c.DBPass, c.DBName, c.DBPort, c.DBSSLMode,
	)
}

func getStr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func getInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return i
}
