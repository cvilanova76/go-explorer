package services

import (
	"context"
	"goexplorer/internal/dtos"
	"goexplorer/internal/models"
	"goexplorer/internal/repositories"
	"log"
	"strings"
)

const (
	MaxNewsArticles = 10
)

type ICountryService interface {
	GetCountryInfo(ctx context.Context, countryName string) (*dtos.CountryAPIResponseDTO, error)
}

type CountryService struct {
	countryRepo repositories.ICountryRepository
	newsRepo    repositories.INewsRepository
}

// NewCountryService constructs a new CountryService instance.
func NewCountryService(countryRepo repositories.ICountryRepository, newsRepo repositories.INewsRepository) *CountryService {
	return &CountryService{
		countryRepo: countryRepo,
		newsRepo:    newsRepo,
	}
}

func (c CountryService) GetCountryInfo(ctx context.Context, countryName string) (*dtos.CountryAPIResponseDTO, error) {
	country, err := c.countryRepo.GetCountryByName(ctx, countryName)
	if err != nil {
		return nil, err
	}

	var newsArticles []models.Article
	var newsTotal int

	countryCode := strings.ToLower(country.CCA2)

	news, err := c.newsRepo.GetNewsByCountry(ctx, countryCode)
	if err != nil {
		log.Printf("Warning: Failed to fetch news for %s: %v", country.Name.Common, err)

		newsArticles = []models.Article{}
		newsTotal = 0
	} else {
		limit := MaxNewsArticles
		if len(news.Articles) < limit {
			limit = len(news.Articles)
		}
		newsArticles = news.Articles[:limit]
		newsTotal = news.TotalArticles
	}

	response := &dtos.CountryAPIResponseDTO{
		Country:      country,
		NewsArticles: newsArticles,
		NewsTotal:    newsTotal,
	}

	return response, nil
}
