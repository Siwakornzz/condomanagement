
# Config Package

The `config` package is responsible for loading and managing application configuration. It loads configuration values from environment variables and a YAML configuration file, and combines them into a structured configuration object that can be used throughout the application.

## Overview

This package provides functionality to load and manage configuration from the following sources:

1. **Environment Variables**: Loaded using `godotenv` and can be used for sensitive data like database credentials, Redis URL, and JWT secret.
2. **YAML Configuration File**: Loaded using `viper`. This file contains the application's configuration settings like app name, version, logging, API settings, etc.

### Configuration Structure

The `Config` struct combines all the application settings. The configuration is structured as follows:

```go
type Config struct {
    // Database
    DBHost string `mapstructure:"DB_HOST"`
    DBPort string `mapstructure:"DB_PORT"`
    DBUser string `mapstructure:"DB_USER"`
    DBPass string `mapstructure:"DB_PASS"`
    DBName string `mapstructure:"DB_NAME"`

    // Redis
    RedisURL string `mapstructure:"REDIS_URL"`

    // JWT Secret
    JwtSecret string `mapstructure:"JWT_SECRET"`

    // App settings (from YAML)
    App AppConfig `mapstructure:"app"`

    // Logging settings (from YAML)
    Logging LoggingConfig `mapstructure:"logging"`

    // API settings (from YAML)
    Api ApiConfig `mapstructure:"api"`
}
```

### Sub-Structs

- `AppConfig`: Contains settings related to the application such as name, version, description, host, and port.
- `LoggingConfig`: Contains logging settings like log level and format.
- `ApiConfig`: Contains API-related settings like rate limit and timeout.

```go
type AppConfig struct {
    Name        string  `mapstructure:"name"`
    Version     float64 `mapstructure:"version"`
    Description string  `mapstructure:"description"`
    Host        string  `mapstructure:"host"`
    Port        int     `mapstructure:"port"`
}

type LoggingConfig struct {
    Level  string `mapstructure:"level"`
    Format string `mapstructure:"format"`
}

type ApiConfig struct {
    RateLimit string `mapstructure:"rate_limit"`
    Timeout   string `mapstructure:"timeout"`
}
```

- The `AppConfig` struct also includes a method `GetAddress()` that returns the full address by combining `Host` and `Port`.

```go
func (a AppConfig) GetAddress() string {
    return fmt.Sprintf("%s:%d", a.Host, a.Port)
}
```

## How to Use

### 1. Loading Configuration

To load the configuration, simply call the `LoadConfig()` function:

```go
config, err := config.LoadConfig()
if err != nil {
    log.Fatalf("Error loading configuration: %v", err)
}
```

This function loads the following:

- Reads the `.env` file (if present) to load environment variables.
- Reads the `app.yaml` file for non-sensitive configuration data.
- Combines values from the environment variables and YAML file into a single structured `Config` object.

### 2. Config File (app.yaml)

The configuration is expected to be stored in a `app.yaml` file in the `../configs` directory. The file should follow this format:

```yaml
app:
  name: condomanagement
  version: 1.0.0
  description: Condo Management System
  host: localhost
  port: 8080

logging:
  level: debug
  format: json

api:
  rate_limit: 100
  timeout: 30s
```

### 3. Environment Variables

Sensitive information such as database credentials, Redis URL, and JWT secret should be stored in a `.env` file. This is an example of what it should look like:

```ini
# Database
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=condo_management

# Redis
REDIS_URL=redis://redis:6379

# JWT Secret
JWT_SECRET=mysecretkey
```

### 4. Merging Configuration

The `viper` package automatically merges the values from the environment variables with the values from the `app.yaml` file. Values from the environment variables will override any matching keys from the YAML file.

For example:

- If `DB_HOST` is provided in the `.env` file, it will override the corresponding value in `app.yaml`.

## Example Usage

After calling `LoadConfig()`, you can access the configuration like this:

```go
fmt.Println("App Name:", config.App.Name)
fmt.Println("Database Host:", config.DBHost)
fmt.Println("API Rate Limit:", config.Api.RateLimit)
fmt.Println("JWT Secret:", config.JwtSecret)
fmt.Println("App Address:", config.App.GetAddress())
```

## Error Handling

- If there is an error while loading the `.env` file, a warning message will be shown, but it will not stop the application from running.
- If the YAML file cannot be found, a warning will be shown, and only the environment variables will be used.

## Additional Notes

- The `.env` file is useful for overriding sensitive settings, such as database credentials and JWT secrets, and should be stored securely.
- The YAML file (`app.yaml`) is typically used for general configuration like application settings, logging, and API configuration.

## Dependencies

- `github.com/joho/godotenv`: For loading environment variables from a `.env` file.
- `github.com/spf13/viper`: For loading and managing configuration from multiple sources (YAML, environment variables).


```
               ___
             _//_\\
           ,"    //".
          /          \
        _/           |
       (.-,--.       |
       /o/  o \     /
       \_\    /  /\/\
       (__`--'   ._)
       /  `-.     |
      (     ,`-.  |
       `-,--\_  ) |-.
        _`.__.'  ,-' \
       |\ )  _.-'    |
       i-\.'\     ,--+.
     .' .'   \,-'/     \
    / /         /       \
    7_|         |       |
    |/          "i.___.j"
    /            |     |\
   /             |     | \
  /              |     |  |
  |              |     |  |
  |____          |     |-i'
   |   """"----""|     | |
   \           ,-'     |/
    `.         `-,     |
     |`-._      / /| |\ \
     |    `-.   `' | ||`-'
     |      |      `-'|
     |      |         |
     |      |         |
     |      |         |
     |      |         |
     |      |         |
     |      |         |
     )`-.___|         |
   .'`-.____)`-.___.-'(
 .'        .'-._____.-i
/        .'           |h
`-------/         .   |j
        `--------' "--'w

```