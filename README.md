# wpscan-api

Go + Echo REST API for WordPress security scanning powered by [WPScan](https://wpscan.com/).

## Prerequisites

- Go 1.23+
- [WPScan](https://github.com/wpscanteam/wpscan) installed and available in PATH

## Getting started

```bash
go run ./cmd/server
```

The server starts on `http://localhost:8080`.

## API endpoints

| Method | Path | Description |
|--------|------|-------------|
| POST | `/api/v1/scans` | Trigger a new scan |
| GET | `/api/v1/scans` | List all scans |
| GET | `/api/v1/scans/:id` | Get scan details |

## Example

```bash
# Start a scan
curl -X POST http://localhost:8080/api/v1/scans \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com"}'

# Check results
curl http://localhost:8080/api/v1/scans/<scan-id>
```
