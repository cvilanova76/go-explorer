# GoExplorer

GoExplorer is a Go-based web application that provides comprehensive country information and news data. The project began as a CLI tool and has been refactored to expose an HTTP API while keeping the CLI as fallback.

## Overview

GoExplorer allows users to retrieve detailed country information (from RestCountries) and current top news stories (from GNews.io).  The choice of APIs reflects the practitioner's personal interest in, and knowledge of, geopolitics and media. The refactor introduced a server-oriented architecture while preserving the original CLI behavior for quick, local queries.

## Roadmap

### Current Features (Phase 1)
- **Country Information**: Integration with RestCountries API for comprehensive country data retrieval
- **News Integration**: Integration with GNews.io API for fetching top news stories by country
- Real-time news data access from multiple sources
- **HTTP Server**: Expose an HTTP API endpoint (e.g. `/api/country/{name}`) to serve country data and news
- **CLI fallback**: Keep a lightweight CLI mode for quick lookups

### Planned Features (subsequent phases)
- **Web Scraper**: Custom news data scraper targeting leading media sources from specific countries
- **Access through authentication**: User authentication to gain full access to country and news database (basic access without login)
- **Database**: Integration with a database to store country and news data
- **Web-based interface**: Development of basic front-end to showcase functionalities

## Implementation Phases

### Phase 1: Server foundation (plus CLI)
- **Primary Focus**: HTTP server exposing a small REST API and CLI for ad-hoc use
- **Learning Objective**: Organize code for server usage (HTTP handlers, DTOs), while keeping clean CLI flow
- **Deliverables**:
  - HTTP server exposing `/api/country/{name}`
  - CLI command for single-country lookups
  - Clean, modular Go codebase following best practices
  - Comprehensive error handling and JSON-friendly responses

**Example Usage:**

Run the server (default port 8080):

```bash
go run ./cmd/cli --serve --port 8080
```

Request from the browser or curl:

```bash
curl http://localhost:8080/api/country/Brazil
```

Quick CLI lookup (positional country name):

```bash
go run ./cmd/cli Brazil
```

### Phase 2: Web Interface (Future Enhancement)
- **Secondary Focus**: Web-based user interface
- **Benefits**: 
  - Rich visual presentation (flags, maps, charts)
  - Better user experience for non-technical users
  - Portfolio showcase enhancement
- **Implementation**: Web app using the same underlying Go modules developed in Phase 1

## Project Goals

The application will enable users to:
- Search for detailed information about any country
- Retrieve current top news stories for a selected country
- Access real-time news data from multiple sources
- (Future) Scrape news from country-specific leading media outlets

## Technology Stack

- **Language**: Go (Golang)
- **APIs**: 
  - [Restcountries.com](https://restcountries.com) - Country information (direct HTTP client implementation)
  - [GNews.io](https://gnews.io) - News data (direct HTTP client implementation)

## Requirements

### Functional Requirements

#### User Input
- The program must accept a country name as a command-line argument
- Support for various country name formats (official name, common name, etc.)

#### API Integration
- The program must connect to the public Restcountries API to fetch information about the user-specified country. Wrapper abandoned due to data limitation
- Integration with GNews.io API for fetching top news stories by country. API key is required.

#### Data Processing
- API responses (in JSON format) must be processed and decoded into Go data structures (structs)
- Efficient handling of nested JSON data from multiple APIs

#### Error Handling
The program must handle errors clearly and gracefully, including:
- Country not found
- API connection problems
- Invalid input arguments
- API rate limiting or authentication issues
- Network connectivity issues

### Technical Requirements
- Clean, modular code structure
- Proper separation of concerns
- Comprehensive error handling
- Clear documentation
- Following Go best practices and conventions

### File organization

```
goexplorer/
├── cmd/
│   └── cli/
│       └── main.go              # CLI / server entry point
├── configs/
│   └── config.go                # Configuration struct and loading logic
├── internal/                    # Private application code (cannot be imported externally)
│   ├── clients/                 # HTTP client wrappers and shared HTTP code
│   ├── dtos/                    # DTOs used to shape API responses
│   ├── handlers/                # HTTP handlers for server endpoints
│   ├── models/                  # Domain models for country, news data and users
│   ├── repositories/            # Repository interfaces and implementations
│   └── services/                # Business logic and orchestration
├── docs/
│   └── README.md                # This documentation file
├── .env                         # Environment variables (API keys, host URLs)
├── .gitattributes               # line-ending normalization
├── Makefile
├── go.mod                       # Go module file
└── go.sum                       # Go module checksums
```

## Getting started

### Prerequisites

- Go 1.19 or higher
- Internet connection for API access

### Installation

1. Clone the repository:
  ```bash
  git clone <repository-url>
  cd goexplorer
  ```

2. Create a `.env` file in the project root (copy from `.env.example` if present) and set required variables (see below).

3. Install dependencies:
  ```bash
  go mod tidy
  ```

4. Run the application

- Start the HTTP server (recommended for the refactor):

```bash
go run ./cmd/cli --serve --port 8080
```

- Or use the CLI fallback for a single country lookup:

```bash
go run ./cmd/cli Brazil
```

Makefile helpers:

```bash
make run COUNTRY=Brazil      # run CLI mode
make serve PORT=8080           # start server (if Makefile target exists)
make build                     # build the binary into ./bin
```

## API configuration

### Required API keys

- **GNews.io**: Sign up at [gnews.io](https://gnews.io) to get your API key
- **RestCountries.com**: No API key required - free public API

### Environment variables

Create a `.env` file in the project root (or copy from `.env.example`):

```env
REST_COUNTRIES_API_BASE_URL=https://restcountries.com/v3.1
GNEWS_API_BASE_URL=https://gnews.io/api/v4
GNEWS_API_KEY=your_gnews_api_key_here
```

## Usage

The HTTP API exposes the endpoint:

```
GET /api/country/{name}
```

Examples below show expected JSON responses for successful and error cases.

### Success response (200)

Request:

```bash
curl -s http://localhost:8080/api/country/Brazil
```

Example JSON response:

```json
{
  "country": {
    "name": {"common": "Brazil", "official": "Federative Republic of Brazil"},
    "capital": ["Brasilia"],
    "region": "Americas",
    "subregion": "South America",
    "population": 212559409,
    "demonyms": {"eng": {"f": "Brazilian", "m": "Brazilian"}},
    "currencies": {"BRL": {"name": "Brazilian real", "symbol": "R$"}},
    "languages": {"por": "Portuguese"},
    "flags": {"png": "https://flagcdn.com/w320/br.png", "svg": "https://flagcdn.com/br.svg"},
    "maps": {"googleMaps": "https://goo.gl/maps/waCKk21HeeqFzkNC9", "openStreetMaps": "https://www.openstreetmap.org/relation/59470"},
    "cca2": "BR"
  },
  "newsArticles": [
    {
      "id": "",
      "title": "Major news headline for Brazil",
      "description": "Short description of the article...",
      "content": "Full article content or excerpt...",
      "url": "https://news.example.com/article/123",
      "image": "https://news.example.com/image.jpg",
      "publishedAt": "2025-10-25T12:34:56Z",
      "lang": "en",
      "source": {"id": "news-example", "name": "News Example", "url": "https://news.example.com", "country": "br"}
    }
  ],
  "newsTotal": 1
}
```

Notes:
- `country` is the `models.Country` object mapped into the response DTO.
- `newsArticles` is a slice of `models.Article` (limited to a configured maximum).
- `newsTotal` is the total number of articles reported by the news provider for the country.

### Error responses

If the country is not found or an internal error occurs, the service returns a JSON error body with an `error` field.

Example 404 / Not Found (country not found):

```http
HTTP/1.1 404 Not Found
Content-Type: application/json

{"error": "country 'Atlantis' not found"}
```

Example 500 / Internal Server Error (upstream failure):

```http
HTTP/1.1 500 Internal Server Error
Content-Type: application/json

{"error": "failed to fetch news: gnews api call failed: ..."}
```

The handler returns meaningful error messages for common failure modes (missing parameter, not found, upstream API errors). Clients should inspect the HTTP status code and the `error` field in the JSON body.

## Contributing

This is a learning project following a mentored development approach. All contributions should align with Go best practices and the project's educational objectives.

## License

[License information to be added]

## Acknowledgments

- [Restcountries.com](https://restcountries.com/) for country data
- [GNews.io](https://gnews.io) for news data
