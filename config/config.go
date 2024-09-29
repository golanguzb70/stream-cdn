package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config is a struct that holds the configuration for the application.
type Config struct {
	// Port is the port on which the application listens.
	Port            int
	OriginServerURL string
	AllowedOrigins  []string
}

// NewConfig creates a new Config with default values.
func New() Config {
	c := Config{}
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	c.Port = cast.ToInt(getOrReturnDefault("PORT", 8081))

	c.OriginServerURL = cast.ToString(getOrReturnDefault("ORIGIN_SERVER_URL", "http://localhost:8080"))
	c.AllowedOrigins = strings.Split(cast.ToString(getOrReturnDefault("ALLOWED_ORIGINS", "http://localhost:8080")), ";")

	if len(c.AllowedOrigins) > 0 {
		if c.AllowedOrigins[0] == "" {
			c.AllowedOrigins = []string{}
		}
	}

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return value
}
