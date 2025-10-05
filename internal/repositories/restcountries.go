package repositories

import (
	"encoding/json"
	"fmt"
	"goexplorer/configs"
	"goexplorer/internal/models"
	"io"
	"net/http"
	"net/url"
	"time"
)

type CountryRepository interface {
	GetCountryByName(name string) (*models.Country, error)
	// GetCountryByCapital(capital string) (*models.Country, error)
	// GetCountryByRegion(region string) ([]models.Country, error)
	// GetCountryByCurrency(currency string) ([]models.Country, error)
	// GetCountryByLanguage(language string) ([]models.Country, error)
	// FilterBy there can be up to 10 filter fields
}

type RestCountriesAPIClient struct {
	Client  *http.Client
	BaseURL string
}

func NewRestCountriesAPIClient(config *configs.Config) *RestCountriesAPIClient {
	return &RestCountriesAPIClient{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
		BaseURL: config.RestCountriesAPIBaseURL,
	}
}

func (c *RestCountriesAPIClient) GetCountryByName(name string) (*models.Country, error) {
	url := fmt.Sprintf("%s/name/%s?fullText=true", c.BaseURL, url.PathEscape(name))

	resp, err := c.Client.Get(url)
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
