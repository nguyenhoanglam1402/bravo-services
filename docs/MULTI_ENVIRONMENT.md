# Multi-Environment Configuration Guide

## Overview

This application supports three different environments: **local**, **dev** (development), and **prod** (production). Each environment has its own configuration file and specific defaults.

## Environment Detection

The application automatically detects the environment based on the following environment variables (in order of priority):

1. `APP_ENV`
2. `ENVIRONMENT`
3. `GO_ENV`

### Supported Values:

- **Local**: `local`
- **Development**: `dev`, `development`
- **Production**: `prod`, `production`

If no environment is specified, it defaults to `local`.

## Environment Files

### Priority Order

The application loads environment files in the following order:

#### Local Environment

1. `.env.local`
2. `.env` (fallback)

#### Development Environment

1. `.env.dev`
2. `.env.development` (fallback)
3. `.env` (fallback)

#### Production Environment

1. `.env.prod`
2. `.env.production` (fallback)
3. `.env` (fallback)

## Environment-Specific Defaults

### Local Environment (APP_ENV=local)

```bash
POSTGRES_HOST=localhost
POSTGRES_USER=postgres
POSTGRES_DB=bravo_db_local
POSTGRES_PORT=5432
PORT=8080
GIN_MODE=debug
JWT_SECRET=local-secret-key
JWT_EXPIRES_IN=24h
```

### Development Environment (APP_ENV=dev)

```bash
POSTGRES_HOST=localhost
POSTGRES_USER=postgres
POSTGRES_DB=bravo_db_dev
POSTGRES_PORT=5432
PORT=8080
GIN_MODE=debug
JWT_SECRET=dev-secret-key
JWT_EXPIRES_IN=24h
```

### Production Environment (APP_ENV=prod)

```bash
POSTGRES_HOST=db
POSTGRES_USER=postgres
POSTGRES_DB=bravo_db
POSTGRES_PORT=5432
PORT=8080
GIN_MODE=release
JWT_SECRET=change-this-in-production
JWT_EXPIRES_IN=1h
```

## Setup Instructions

### 1. Choose Your Environment

Set the environment variable:

```bash
export APP_ENV=local    # or dev, prod
```

### 2. Create Environment File

Copy the appropriate template:

```bash
# For local development
cp .env.example .env.local

# For development environment
cp .env.example .env.dev

# For production environment
cp .env.example .env.prod
```

### 3. Configure Your Environment

Edit the appropriate `.env.{environment}` file with your specific values.

## Usage in Code

### Check Current Environment

```go
import "bravo-service/api/internal/config"
import "bravo-service/api/packages/utils"

// Using config package
cfg := config.GetConfig()
currentEnv := cfg.Environment

// Using utility functions
isLocal := config.IsLocal()
isDev := config.IsDevelopment()
isProd := config.IsProduction()

// Using utils package
currentEnv := utils.GetCurrentEnvironment()
isLocal := utils.IsLocal()
isDev := utils.IsDevelopment()
isProd := utils.IsProduction()
```

### Environment-Specific Logic

```go
if config.IsProduction() {
    // Production-specific code
    log.SetLevel(log.InfoLevel)
} else {
    // Development/Local specific code
    log.SetLevel(log.DebugLevel)
}
```

## Running the Application

### Local Environment

```bash
# Set environment (optional, defaults to local)
export APP_ENV=local
go run main.go
```

### Development Environment

```bash
export APP_ENV=dev
go run main.go
```

### Production Environment

```bash
export APP_ENV=prod
go run main.go
```

## Docker Support

### Using Environment Variables

```bash
docker run -e APP_ENV=prod -e POSTGRES_PASSWORD=secret your-app
```

### Using Environment Files

```bash
# Development
docker run --env-file .env.dev your-app

# Production
docker run --env-file .env.prod your-app
```

### Docker Compose

```yaml
# docker-compose.dev.yml
services:
  app:
    environment:
      - APP_ENV=dev
    env_file:
      - .env.dev

# docker-compose.prod.yml
services:
  app:
    environment:
      - APP_ENV=prod
    env_file:
      - .env.prod
```

## Best Practices

1. **Never commit environment files** to version control (they're in `.gitignore`)
2. **Use strong, unique secrets** for each environment
3. **Keep production secrets secure** and use secret management systems
4. **Test with different environments** before deploying
5. **Use environment-specific database names** to avoid conflicts
6. **Configure different ports** for different environments when running locally

## Environment Variables Summary

| Variable         | Local Default  | Dev Default  | Prod Default | Description            |
| ---------------- | -------------- | ------------ | ------------ | ---------------------- |
| `APP_ENV`        | local          | dev          | prod         | Environment identifier |
| `POSTGRES_HOST`  | localhost      | localhost    | db           | Database host          |
| `POSTGRES_DB`    | bravo_db_local | bravo_db_dev | bravo_db     | Database name          |
| `PORT`           | 8080           | 8081         | 8080         | Server port            |
| `GIN_MODE`       | debug          | debug        | release      | Gin framework mode     |
| `JWT_EXPIRES_IN` | 24h            | 24h          | 1h           | JWT token expiration   |

## Troubleshooting

### Environment Not Detected

- Check that `APP_ENV` is set correctly
- Verify the environment file exists
- Check file permissions

### Database Connection Issues

- Verify database configuration for your environment
- Check if the database exists
- Confirm credentials are correct

### Configuration Not Loading

- Check file paths and naming
- Verify environment file syntax
- Look for application logs for loading messages
