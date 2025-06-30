# URL Shortener - Go Backend Project

A high-performance URL shortening service built with Go and Gin framework. 

##  Features

- **URL Shortening**: Convert long URLs to short, memorable codes
- **Analytics Tracking**: Track click counts for each shortened URL
- **Fast Redirects**: Sub-millisecond redirect performance
- **RESTful API**: Clean, well-documented API endpoints
- **In-Memory Storage**: Fast access with automatic cleanup
- **Error Handling**: Comprehensive validation and error responses
- **Health Monitoring**: Built-in health check endpoint

## Tech Stack

- **Language**: Go 1.24+
- **Framework**: Gin (HTTP web framework)
- **Storage**: In-memory (can be extended to PostgreSQL/Redis)
- **Validation**: Gin's built-in validation
- **Testing**: Built-in Go testing

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Health check endpoint |
| `POST` | `/shorten` | Create a shortened URL |
| `GET` | `/:code` | Redirect to original URL |
| `GET` | `/analytics/:code` | Get analytics for a URL |

## 🚀 Quick Start

### Prerequisites

- Go 1.24 or higher
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <your-repo-url>
   cd url-shortener
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Run the server**
   ```bash
   go run cmd/server/main.go
   ```

4. **Test the service**
   ```bash
   curl http://localhost:8080/health
   ```

## 📖 Usage Examples

### 1. Create a Shortened URL

```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://www.google.com"}'
```

**Response:**
```json
{
  "code": "ABC123",
  "original_url": "https://www.google.com",
  "short_url": "localhost:8080/ABC123"
}
```

### 2. Redirect to Original URL

```bash
curl -I http://localhost:8080/ABC123
```

**Response:**
```
HTTP/1.1 301 Moved Permanently
Location: https://www.google.com
```

### 3. Get Analytics

```bash
curl http://localhost:8080/analytics/ABC123
```

**Response:**
```json
{
  "code": "ABC123",
  "original_url": "https://www.google.com",
  "clicks": 5,
  "short_url": "localhost:8080/ABC123"
}
```

### 4. Health Check

```bash
curl http://localhost:8080/health
```

**Response:**
```json
{
  "status": "ok",
  "service": "url-shortener",
  "version": "1.0.0",
  "mode": "in-memory"
}
```

##  Project Structure

```
url-shortener/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── internal/
│   ├── handlers/
│   │   ├── url.go          # URL shortening and redirect handlers
│   │   └── analytics.go    # Analytics handler
│   └── database/           # Database interfaces (for future use)
├── go.mod                  # Go module file
├── go.sum                  # Go module checksums
└── README.md              # This file
```

##  Configuration

The service can be configured using environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `GIN_MODE` | `debug` | Gin mode (debug/release) |

##  Testing

Run the tests:

```bash
go test ./...
```

##  Production Considerations

For production deployment, consider:

1. **Database**: Replace in-memory storage with PostgreSQL
2. **Caching**: Add Redis for fast lookups
3. **Rate Limiting**: Implement request rate limiting
4. **Monitoring**: Add Prometheus metrics and Grafana dashboards
5. **Containerization**: Use Docker for deployment
6. **Load Balancing**: Use multiple instances behind a load balancer
7. **SSL/TLS**: Add HTTPS support
8. **Logging**: Implement structured logging


