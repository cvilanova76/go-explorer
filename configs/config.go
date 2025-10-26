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
	RestCountriesAPIHostURL string
	GNewsAPIHostURL         string
	GNewsAPIKey             string
}

func LoadConfig() *Config {

	_, b, _, _ := runtime.Caller(0)

	basePath := filepath.Dir(b)

	envPath := filepath.Join(basePath, "..", ".env")

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	return &Config{
		Environment:             os.Getenv("GO_ENV"),
		RestCountriesAPIHostURL: os.Getenv("REST_COUNTRIES_API_BASE_URL"),
		GNewsAPIHostURL:         os.Getenv("GNEWS_API_BASE_URL"),
		GNewsAPIKey:             os.Getenv("GNEWS_API_KEY"),
	}
}
