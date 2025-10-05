package services

import (
	"fmt"
	"goexplorer/internal/repositories"
	"strings"
)

type ICountryService interface {
	DisplayCountryByName(countryName string) error
}
type CountryService struct {
	repo repositories.CountryRepository
}

func NewCountryService(repo repositories.CountryRepository) ICountryService {
	return &CountryService{repo: repo}
}

func (c CountryService) DisplayCountryByName(countryName string) error {
	country, err := c.repo.GetCountryByName(countryName)
	if err != nil {
		return err
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

	return nil
}
