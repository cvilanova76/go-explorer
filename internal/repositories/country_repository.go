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
	"net/url"
)

type ICountryRepository interface {
	GetCountryByName(ctx context.Context, name string) (*models.Country, error)
	// GetCountryByCapital(capital string) (*models.Country, error)
	// GetCountryByRegion(region string) ([]models.Country, error)
	// GetCountryByCurrency(currency string) ([]models.Country, error)
	// GetCountryByLanguage(language string) ([]models.Country, error)
	// FilterBy there can be up to 10 filter fields
}

type RestCountriesAPIClient struct {
	Client  *http.Client
	HostURL string
}

func NewRestCountriesAPIClient(config *configs.Config) *RestCountriesAPIClient {
	return &RestCountriesAPIClient{
		Client:  clients.NewDefaultHTTPClient(),
		HostURL: config.RestCountriesAPIHostURL,
	}
}

func (c RestCountriesAPIClient) GetCountryByName(ctx context.Context, name string) (*models.Country, error) {
	url := fmt.Sprintf("%s/name/%s?fullText=true", c.HostURL, url.PathEscape(name))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("rest countries api call failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("rest countries api call failed: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var countries []models.Country
	err = json.Unmarshal(body, &countries)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	if len(countries) == 0 {
		return nil, fmt.Errorf("country '%s' not found", name)
	}

	return &countries[0], nil

}
