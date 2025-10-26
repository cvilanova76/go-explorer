package dtos

import "goexplorer/internal/models"

type CountryAPIResponseDTO struct {
	Country      *models.Country  `json:"country"`
	NewsArticles []models.Article `json:"newsArticles"`
	NewsTotal    int              `json:"newsTotal"`
}
