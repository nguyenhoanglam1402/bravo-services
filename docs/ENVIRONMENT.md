# Environment Configuration Guide

## Overview

This application uses environment variables for configuration management. The configuration is handled through a centralized config package that loads variables from `.env` files or system environment variables.

## Setup

### 1. Copy the example environment file

```bash
cp .env.example .env
```

### 2. Edit the `.env` file with your actual values

```bash
# Database Configuration
POSTGRES_HOST=localhost
POSTGRES_USER=postgres
POSTGRES_PASSWORD=your_actual_password
POSTGRES_DB=bravo_db
POSTGRES_PORT=5432

# Server Configuration
PORT=8080
GIN_MODE=debug  # or "release" for production

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-here
JWT_EXPIRES_IN=24h
```

## Environment Variables

### Database

- `POSTGRES_HOST`: Database host (default: localhost)
- `POSTGRES_USER`: Database username (default: postgres)
- `POSTGRES_PASSWORD`: Database password (default: empty)
- `POSTGRES_DB`: Database name (default: bravo_db)
- `POSTGRES_PORT`: Database port (default: 5432)

### Server

- `PORT`: Server port (default: 8080)
- `GIN_MODE`: Gin framework mode - "debug" or "release" (default: debug)

### JWT

- `JWT_SECRET`: Secret key for JWT token signing (default: your-secret-key)
- `JWT_EXPIRES_IN`: Token expiration time (default: 24h)

## Usage in Code

### Using the config package

```go
import "bravo-service/api/internal/config"

func someFunction() {
    cfg := config.GetConfig()
    dbHost := cfg.Database.Host
    serverPort := cfg.Server.Port
}
```

### Using utility functions

```go
import "bravo-service/api/packages/utils"

func someFunction() {
    jwtSecret := utils.GetJWTSecret()
    isProduction := utils.IsProduction()
}
```

### Getting environment variables directly

```go
import "bravo-service/api/packages/utils"

func someFunction() {
    customValue := utils.GetEnvOrDefault("CUSTOM_VAR", "default_value")
}
```

## Development vs Production

### Development

Set `GIN_MODE=debug` in your `.env` file for detailed logging and debug features.

### Production

Set `GIN_MODE=release` for optimized performance and minimal logging.

## Security Notes

1. **Never commit the `.env` file** to version control
2. **Use strong, unique JWT secrets** in production
3. **Use secure database passwords**
4. **Consider using environment-specific files** (`.env.development`, `.env.production`)

## Docker Support

For Docker deployments, you can pass environment variables directly:

```bash
docker run -e POSTGRES_HOST=db -e POSTGRES_PASSWORD=secret your-app
```

Or use an environment file:

```bash
docker run --env-file .env your-app
```
