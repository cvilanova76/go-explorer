package main

import (
	"context"
	"flag"
	"fmt"
	"goexplorer/configs"
	"goexplorer/internal/handlers"
	"goexplorer/internal/repositories"
	"goexplorer/internal/services"
	"log"
	"net/http"
	"strings"
)

func main() {
	config := configs.LoadConfig()

	// CLI flags
	serve := flag.Bool("serve", false, "run as HTTP server")
	port := flag.Int("port", 8080, "port for HTTP server")
	flag.Parse()

	// If running as server, start HTTP server and expose handlers
	countryRepo := repositories.NewRestCountriesAPIClient(config)
	newsRepo := repositories.NewNewsAPIClient(config)
	countryService := services.NewCountryService(countryRepo, newsRepo)

	if *serve {
		h := handlers.NewCountryHandler(countryService)
		http.HandleFunc("/api/country/", h.GetCountryInfo)
		addr := fmt.Sprintf(":%d", *port)
		log.Printf("Starting HTTP server on %s", addr)
		log.Fatal(http.ListenAndServe(addr, nil))
		return
	}

	// CLI mode: expect a country name as positional argument
	if flag.NArg() < 1 {
		fmt.Println("Usage: go run . [--serve] <country_name>")
		return
	}

	countryName := strings.TrimSpace(flag.Arg(0))

	// Use a background context for CLI usage
	ctx := context.Background()

	resp, err := countryService.GetCountryInfo(ctx, countryName)
	if err != nil {
		log.Printf("Error while getting country: %v\n", err)
		return
	}

	// Format country information
	fmt.Printf("\nCountry Information:\n")
	fmt.Printf("Name: %s (%s)\n", resp.Country.Name.Common, resp.Country.Name.Official)
	fmt.Printf("Capital: %s\n", strings.Join(resp.Country.Capital, ", "))
	fmt.Printf("Region: %s (%s)\n", resp.Country.Region, resp.Country.Subregion)
	fmt.Printf("Population: %d\n", resp.Country.Population)

	currencies := make([]string, 0)
	for code, currency := range resp.Country.Currencies {
		currencies = append(currencies, fmt.Sprintf("%s (%s %s)", code, currency.Name, currency.Symbol))
	}
	fmt.Printf("Currencies: %s\n", strings.Join(currencies, ", "))

	languages := make([]string, 0)
	for _, lang := range resp.Country.Languages {
		languages = append(languages, lang)
	}
	fmt.Printf("Languages: %s\n", strings.Join(languages, ", "))
	fmt.Printf("Maps: %s\n", resp.Country.Maps.GoogleMaps)

	// Print news articles
	fmt.Printf("\nRecent News:\n")
	for _, article := range resp.NewsArticles {
		fmt.Printf("\nTitle: %s\n", article.Title)
		fmt.Printf("Description: %s\n", article.Description)
		fmt.Printf("URL: %s\n", article.Url)
		fmt.Printf("Published: %s\n", article.PublishedAt)
	}

}
