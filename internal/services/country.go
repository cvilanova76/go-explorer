package services

import (
	"fmt"
	"goexplorer/internal/models"
	"goexplorer/internal/repositories"
	"log"
	"strings"
)

type ICountryService interface {
	DisplayCountryByName(countryName string) error
}

type CountryService struct {
	repo     repositories.CountryRepository
	newsRepo repositories.NewsRepository
}

func NewCountryService(repo repositories.CountryRepository, newsRepo repositories.NewsRepository) ICountryService {
	return &CountryService{repo: repo, newsRepo: newsRepo}
}

func (c CountryService) DisplayCountryByName(countryName string) error {
	country, err := c.repo.GetCountryByName(countryName)
	if err != nil {
		return err
	}

	countryCode := strings.ToLower(strings.TrimSpace(country.CCA2))
	var news *models.News
	if countryCode == "" {
		log.Printf("Warning: No CCA2 code for %s; News section will be empty", country.Name.Common)
		news = &models.News{}
	} else {
		news, err = c.newsRepo.GetNewsByCountry(countryCode)
		if err != nil {
			log.Printf("Warning: Failed to fetch news for %s (%s). News section will be empty: %v",
				country.Name.Common, countryCode, err)
			news = &models.News{}
		}
	}

	fmt.Println("--- Country ---")
	fmt.Printf("Name: %s (%s)\n", country.Name.Common, country.Name.Official)

	fmt.Printf("Capital: ")
	if len(country.Capital) > 0 {
		fmt.Printf("%s\n", strings.Join(country.Capital, ", "))
	} else {
		fmt.Println("N/A")
	}

	fmt.Printf("Region: %s\n", country.Region)
	fmt.Printf("Subregion: %s\n", country.Subregion)
	fmt.Printf("Population: %d\n", country.Population)
	fmt.Printf("Flags: PNG=%s, SVG=%s\n", country.Flags.PNG, country.Flags.SVG)
	fmt.Printf("Maps: GoogleMaps=%s, OpenStreetMaps=%s\n", country.Maps.GoogleMaps, country.Maps.OpenStreetMaps)
	fmt.Printf("CCA2: %s\n", country.CCA2)

	fmt.Printf("Demonym: ")
	if engDemon, ok := country.Demonyms.Eng["m"]; ok {
		fmt.Printf("%s\n", engDemon)
	} else {
		fmt.Println("N/A")
	}

	fmt.Printf("Currencies: ")
	var currencies []string
	for _, currency := range country.Currencies {
		currencies = append(currencies, fmt.Sprintf("%s (%s)", currency.Name, currency.Symbol))
	}
	fmt.Println(strings.Join(currencies, ", ")) //para concatenar

	fmt.Printf("Languages: ")
	var languages []string
	for code, language := range country.Languages {
		languages = append(languages, fmt.Sprintf("%s (%s)", language, code))
	}
	fmt.Println(strings.Join(languages, ", "))

	fmt.Println("\n--- News Articles ---")
	if len(news.Articles) == 0 {
		fmt.Println("No articles found.")
	} else {
		for i, article := range news.Articles {
			fmt.Printf("\nArticle %d:\n", i+1)
			fmt.Printf("Title: %s\n", article.Title)
			fmt.Printf("Description: %s\n", article.Description)
			fmt.Printf("URL: %s\n", article.Url)
			fmt.Printf("Published At: %s\n", article.PublishedAt.Format("2006-01-02 15:04:05"))
		}
	}

	return nil
}
