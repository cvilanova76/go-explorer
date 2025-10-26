package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"goexplorer/configs"
	"goexplorer/internal/clients"
	"goexplorer/internal/models"
	"io"
	"net/http"
)

type INewsRepository interface {
	GetNewsByCountry(ctx context.Context, country string) (*models.News, error)
	//GetNewsByCategory(category string) (*models.News, error)
}
type GNewsAPIClient struct {
	Client  *http.Client
	HostURL string
	APIKey  string
}

func NewNewsAPIClient(config *configs.Config) *GNewsAPIClient {
	return &GNewsAPIClient{
		Client:  clients.NewDefaultHTTPClient(),
		HostURL: config.GNewsAPIHostURL,
		APIKey:  config.GNewsAPIKey,
	}
}

func (g *GNewsAPIClient) GetNewsByCountry(ctx context.Context, country string) (*models.News, error) {
	url := fmt.Sprintf("%s/top-headlines?country=%s&apikey=%s", g.HostURL, country, g.APIKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := g.Client.Do(req)
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
