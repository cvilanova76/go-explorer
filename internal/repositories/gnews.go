package repositories

import (
	"encoding/json"
	"fmt"
	"goexplorer/configs"
	"goexplorer/internal/models"
	"io"
	"net/http"
	"time"
)

type NewsRepository interface {
	GetNewsByCountry(country string) (*models.News, error)
	//GetNewsByCategory(category string) (*models.News, error)
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
func (g *GNewsAPIClient) GetNewsByCountry(country string) (*models.News, error) {
	url := fmt.Sprintf("%s/top-headlines?country=%s&apikey=%s", g.BaseURL, country, g.APIKey)

	resp, err := g.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("gnews api call failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("gnews api call failed: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var news models.News
	err = json.Unmarshal(body, &news)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	if len(news.Articles) == 0 {
		return nil, fmt.Errorf("no news found for country '%s'", country)
	}

	return &news, nil

}
