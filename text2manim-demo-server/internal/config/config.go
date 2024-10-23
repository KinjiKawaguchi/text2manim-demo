package config

import (
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	ServerPort  string

	SupabaseHost     string
	SupabaseUser     string
	SupabasePassword string
	SupabaseDBName   string
	SupabasePort     string

	Text2manimApiEndpoint string
	Text2manimApiKey      string

	RateLimitRequests int
	RateLimitInterval time.Duration

	AllowedOrigins []string
}

func Load(logger *slog.Logger) *Config {
	err := godotenv.Load()
	if err != nil {
		logger.Warn("Error loading .env file")
	}

	rateLimitRequests, err := strconv.Atoi(getEnv("RATE_LIMIT_REQUESTS", "100"))
	if err != nil {
		logger.Warn("Invalid RATE_LIMIT_REQUESTS, using default", "default", 100)
		rateLimitRequests = 100
	}

	rateLimitInterval, err := strconv.Atoi(getEnv("RATE_LIMIT_INTERVAL", "3600"))
	if err != nil {
		logger.Warn("Invalid RATE_LIMIT_INTERVAL, using default", "default", 3600)
		rateLimitInterval = 3600
	}

	config := &Config{
		Environment: getEnv("ENVIRONMENT", "development"),
		ServerPort:  getEnv("SERVER_PORT", "8080"),

		SupabaseHost:     getEnv("SUPABASE_HOST", ""),
		SupabaseUser:     getEnv("SUPABASE_USER", ""),
		SupabasePassword: getEnv("SUPABASE_PASSWORD", ""),
		SupabaseDBName:   getEnv("SUPABASE_DBNAME", ""),
		SupabasePort:     getEnv("SUPABASE_PORT", ""),

		Text2manimApiEndpoint: getEnv("TEXT2MANIM_API_ENDPOINT", ""),
		Text2manimApiKey:      getEnv("TEXT2MANIM_API_KEY", ""),

		RateLimitRequests: rateLimitRequests,
		RateLimitInterval: time.Duration(rateLimitInterval) * time.Second,

		AllowedOrigins: strings.Split(getEnv("ALLOWED_ORIGINS", "http://localhost:3000"), ","),
	}

	logger.Info("Configuration loaded successfully")
	return config
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
