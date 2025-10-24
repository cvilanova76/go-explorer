package main

import (
	"fmt"
	"goexplorer/configs"
	"goexplorer/internal/repositories"
	"goexplorer/internal/services"
	"log"
	"os"
	"strings"
)

func main() {
	config := configs.LoadConfig()

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <country_name>")
		return
	}

	countryName := strings.TrimSpace(os.Args[1])

	repo := repositories.NewRestCountriesAPIClient(config)
	newsRepo := repositories.NewGNewsAPIClient(config)
	s := services.NewCountryService(repo, newsRepo)

	err := s.DisplayCountryByName(countryName)
	if err != nil {
		log.Printf("Error while getting country: %v\n", err)
		return
	}
}
