package configs

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment             string
	RestCountriesAPIBaseURL string
	GNewsAPIBaseURL         string
	GNewsAPIKey             string
}

func LoadConfig() *Config {

	_, b, _, _ := runtime.Caller(0)

	basePath := filepath.Dir(b)

	envPath := filepath.Join(basePath, "..", ".env")

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Failed to load .env file: %v", err)
	}

	return &Config{
		Environment:             os.Getenv("GO_ENV"),
		RestCountriesAPIBaseURL: os.Getenv("REST_COUNTRIES_API_BASE_URL"),
		GNewsAPIBaseURL:         os.Getenv("GNEWS_API_BASE_URL"),
		GNewsAPIKey:             os.Getenv("GNEWS_API_KEY"),
	}
}
