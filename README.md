# GoExplorer

GoExplorer is a Go-based web application that provides country information and related news. The project started as a CLI tool and has been refactored to expose an HTTP API while keeping the CLI mode as fallback.

## Quick start

1. Create a `.env` file in the project root (copy from `.env.example` if present) and set the required API keys and base URLs:

```env
REST_COUNTRIES_API_BASE_URL=https://restcountries.com/v3.1
GNEWS_API_BASE_URL=https://gnews.io/api/v4
GNEWS_API_KEY=your_gnews_api_key_here
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run the application

- Start the HTTP server (default port 8080):

```bash
go run ./cmd/cli --serve --port 8080
```

The server exposes a simple endpoint to get country info and related news:

```
GET /api/country/{name}
```

- Use the CLI mode for quick lookups (positional country name):

```bash
go run ./cmd/cli Brazil
```

- Makefile helpers (from project root):

```bash
make run COUNTRY=Brazil      # run CLI mode
make serve PORT=8080           # (if you want) start the server via Makefile
make build                     # build the binary into ./bin
```

## Documentation

📖 **Full documentation is available in [`docs/README.md`](docs/README.md)**

The complete documentation includes:
- Detailed project overview and goals
- Implementation phases and architecture
- API configuration and usage
- Code structure and development guidelines

## Project structure

```
goexplorer/
├── cmd/cli/               # CLI / server entry point (main.go)
├── internal/              # Private application code
│   ├── clients/           # HTTP client wrappers and shared HTTP code
│   ├── dtos/              # Data transfer objects for API responses
│   ├── handlers/          # HTTP handlers (server-side)
│   ├── models/            # Domain models
│   ├── repositories/      # Repository abstractions / API clients
│   └── services/          # Business logic and orchestration
├── configs/               # Configuration loading and env helpers
├── docs/                  # Documentation
├── Makefile               # helper targets (run, serve, build)
└── .gitattributes         # line-ending rules
```

## Contributing

This is a learning project following Go best practices. See [`docs/README.md`](docs/README.md) for detailed development guidelines.
