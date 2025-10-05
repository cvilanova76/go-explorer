# GoExplorer

A Go-based CLI tool application that provides comprehensive country information and news data.

## Overview

GoExplorer is a command-line interface tool that enables users to search for detailed information about any country and retrieve current top news stories for that country. This project demonstrates Go programming proficiency while integrating with multiple external APIs.
The choice of APIs reflects the practitioner's personal interest in, and knowledge of, geopolitics and media.

## Roadmap

### Current Features (Phase 1)
- **Country Information**: Integration with RestCountries API for comprehensive country data retrieval
- **News Integration**: Integration with GNews.io API for fetching top news stories by country
- Real-time news data access from multiple sources
- **Database**: Integration with a database to store country and news data


### Planned Features (subsequent phases)
- **Web Scraper**: Custom news data scraper targeting leading media sources from specific countries
- **Access through authentication**: User authentication to gain full access to country and news database (basic access without login)
- **Web-based interface**: Build web interface that leverages the CLI tool's core Go packages

## Implementation Phases

### Phase 1: CLI Tool Foundation
- **Primary Focus**: Command-line interface application
- **Learning Objective**: Master Go fundamentals (HTTP clients, JSON handling, error management, CLI parsing)
- **Deliverables**:
  - Fully functional CLI tool for country information and news retrieval
  - Clean, modular Go codebase following best practices
  - Comprehensive error handling and user feedback
  - Single binary deployment

**Example Usage:**
```bash
# Basic country lookup
./goexplorer brazil

# Advanced usage with flags
./goexplorer --country="United States" --news-limit=5
./goexplorer --country="Japan" --format=json
```

### Phase 2: Optional Web Interface (Future Enhancement)
- **Secondary Focus**: Web-based user interface
- **Approach**: Build web interface that leverages the CLI tool's core Go packages
- **Benefits**: 
  - Rich visual presentation (flags, maps, charts)
  - Better user experience for non-technical users
  - Portfolio showcase enhancement
- **Implementation**: Web app using the same underlying Go modules developed in Phase 1

**Rationale**: Starting with CLI ensures solid Go foundations before adding UI complexity. The modular design allows easy transition to web interface later.

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
- The program must connect to the public Restcountries API to fetch information about the user-specified country. Wrapper
abandoned due to data limitation
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

#### Output
The program must display the following country information in an organized and easy-to-read format in the terminal:
- Common and Official country name
- Capital city
- Population
- Region
- Subregion
- Languages
- Currencies
- Demonym
- Flag (links to PNG and SVG images)
- Maps (Google Maps and OpenStreetMaps links)
- Top news stories for the country

### Technical Requirements
- Clean, modular code structure
- Proper separation of concerns
- Comprehensive error handling
- Clear documentation
- Following Go best practices and conventions

## Code Structure

To maintain clean and organized code, we follow a modular structure, separating different responsibilities into functions and, when necessary, into separate files.

### File Organization
(Currently)

```
goexplorer/
├── cmd/
│   └── cli/
│       └── main.go              # CLI application entry point
├── configs/
│   └── config.go                # Configuration struct and loading logic
├── internal/                    # Private application code (cannot be imported externally)
│   ├── models/
│   │   ├── country.go           # Domain models for country data
│   │   └── news.go              # Domain models for news data
│   ├── repositories/
│   │   ├── restcountries.go     # RestCountries API client implementation
│   │   └── gnews.go             # GNews API client implementation
│   └── services/
│       └── country.go           # Country service (orchestrates repositories)
├── docs/
│   └── README.md                # This documentation file
├── .env                         # Environment variables (API keys, base URLs)
├── go.mod                       # Go module file
└── go.sum                       # Go module checksums
```

## Getting Started

### Prerequisites

- Go 1.19 or higher
- Internet connection for API access

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd goexplorer
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

## API Configuration

### Required API Keys

- **GNews.io**: Sign up at [gnews.io](https://gnews.io) to get your API key
- **RestCountries.com**: No API key required - free public API

### Environment Variables

Create a `.env` file in the project root:
```
REST_COUNTRIES_API_BASE_URL=https://restcountries.com/v3.1
GNEWS_API_BASE_URL=https://gnews.io/api/v4
GNEWS_API_KEY=your_gnews_api_key_here
```

## Usage

[Usage examples will be added as features are implemented]

## Contributing

This is a learning project following a mentored development approach. All contributions should align with Go best practices and the project's educational objectives.

## License

[License information to be added]

## Acknowledgments

- [Restcountries.com](https://restcountries.com/) for country data
- [GNews.io](https://gnews.io) for news data
