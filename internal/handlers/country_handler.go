package handlers

import (
	"encoding/json"
	"goexplorer/internal/services"
	"log"
	"net/http"
	"strings"
)

type CountryHandler struct {
	CountryService services.ICountryService
}

func NewCountryHandler(countryService services.ICountryService) *CountryHandler {
	return &CountryHandler{
		CountryService: countryService,
	}
}

func JSONError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (h *CountryHandler) GetCountryInfo(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 || pathParts[3] == "" {
		JSONError(w, "Missing country name in URL path. Usage: /api/country/{name}", http.StatusBadRequest)
		return
	}
	countryName := pathParts[3]

	respDto, err := h.CountryService.GetCountryInfo(ctx, countryName)
	if err != nil {
		log.Printf("Error processing request for %s: %v", countryName, err)
		JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respDto)

}
