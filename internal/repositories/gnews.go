package repositories

import (
	"goexplorer/configs"
	"goexplorer/internal/models"
	"net/http"
	"time"
)

type NewsRepository interface {
	GetNewsByCountry(country string) ([]models.News, error)
	//GetNewsByCategory(category string) ([]models.News, error)
}

type GNewsAPIClient struct {
	Client  *http.Client
	BaseURL string
	APIKey  string
}

func NewGNewsAPIClient(config *configs.Config) *GNewsAPIClient {
	return &GNewsAPIClient{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
		BaseURL: config.GNewsAPIBaseURL,
		APIKey:  config.GNewsAPIKey,
	}
}
