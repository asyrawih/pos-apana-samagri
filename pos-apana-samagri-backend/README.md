# POS Apana Samagri Backend

This is the backend service for the POS Apana Samagri system, built with Go and Gin framework.

## Features

- RESTful API for POS operations
- JWT-based authentication
- PostgreSQL database integration
- Structured logging with Zap
- Input validation
- Unit and integration tests

## Project Structure

```
.
├── cmd/
│   └── main.go               # Application entry point
├── internal/
│   ├── config/              # Configuration management
│   ├── handlers/            # HTTP request handlers
│   ├── models/              # Data models
│   ├── repository/          # Database operations
│   ├── service/             # Business logic
│   ├── ai/                  # AI integration
│   └── middleware/          # HTTP middleware
├── pkg/
│   ├── logger/              # Logging utilities
│   └── utils/               # Common utilities
├── tests/                   # Test files
└── swagger/                 # API documentation
```

## Getting Started

### Prerequisites

- Go 1.22 or later
- PostgreSQL 13 or later
- Git

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/pos-apana-samagri-backend.git
   cd pos-apana-samagri-backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. Run the application:
   ```bash
   go run cmd/main.go
   ```

### Environment Variables

Create a `.env` file in the root directory with the following variables:

```
# Server
PORT=8080

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=pos_db
DB_SSLMODE=disable

# JWT
JWT_SECRET=your-secret-key
```

## API Documentation

API documentation is available at `/swagger/index.html` when running the application.

## Testing

Run tests:

```bash
go test -v ./...
```

## Deployment

### Build

```bash
go build -o bin/pos-apana-samagri cmd/main.go
```

### Docker

Build the Docker image:

```bash
docker build -t pos-apana-samagri .
```

Run the container:

```bash
docker run -p 8080:8080 --env-file .env pos-apana-samagri
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
