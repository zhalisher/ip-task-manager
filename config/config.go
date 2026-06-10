package config

import (
	"os"
	"strconv"
)

type Config struct {
	PostgresURI string

	JWTSecret        string
	JWTAccessExpMin  int
	JWTRefreshExpDay int
}

func Load() *Config {
	return &Config{
		PostgresURI:      getEnv("POSTGRES_URI", ""),
		JWTSecret:        getEnv("JWT_SECRET", ""),
		JWTAccessExpMin:  getEnvInt("JWT_ACCESS_EXP_MIN", 15),
		JWTRefreshExpDay: getEnvInt("JWT_REFRESH_EXP_DAY", 7),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
func getEnvInt(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return fallback
}
