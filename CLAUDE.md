# wpscan-api — Instructions

## Project

REST API for WordPress security scanning. Wraps WPScan CLI and exposes results via HTTP endpoints.

## Stack

- Go + Echo v4
- In-memory storage (to be replaced with a database)
- WPScan executed as a subprocess via `os/exec`

## Structure

- `cmd/server/` — Entry point
- `internal/handler/` — Echo route handlers
- `internal/scanner/` — WPScan subprocess integration
- `internal/model/` — Shared data types

## Running

```bash
go run ./cmd/server
```

Server starts on port 8080.
