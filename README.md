# GoExplorer

A Go-based CLI tool for country information and news data.

## Quick Start

```bash
# Get API key from gnews.io
# Set up environment variables in .env file (copy .env.example -> .env)

# Install dependencies
go mod tidy

# Run the application from project root (recommended):
go run ./cmd/cli <country_name>

# Or enter the command directory and run:
cd cmd/cli
go run . <country_name>

# Using Makefile (from project root):
make run COUNTRY=Portugal
```

## Documentation

📖 **Full documentation is available in [`docs/README.md`](docs/README.md)**

The complete documentation includes:
- Detailed project overview and goals
- Implementation phases and architecture
- API configuration and usage
- Code structure and development guidelines

## Project Structure

```
goexplorer/
├── cmd/cli/           # CLI application entry point
├── internal/          # Private application code
│   ├── models/        # Domain models
│   ├── repositories/  # API clients
│   └── services/      # Business logic
├── configs/           # Configuration management
└── docs/              # 📖 Full documentation
```

## Contributing

This is a learning project following Go best practices. See [`docs/README.md`](docs/README.md) for detailed development guidelines.
