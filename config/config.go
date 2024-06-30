package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const ()

// Config ...
type Config struct {
	Environment   string // develop, staging, production
	RedisHost     string
	RedisPort     string
	RedisPassword string

	UserBranchHost string
	UserBranchPort string

	ScheduleBranchHost string
	ScheduleBranchPort string

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "prod"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":1239"))

	c.RedisHost = "localhost"
	c.RedisPort = "6379"
	c.RedisPassword = cast.ToString(getOrReturnDefault("REDIS_PASSWORD", "Amir2414"))

	c.UserBranchHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost"))
	c.UserBranchPort = cast.ToString(getOrReturnDefault("USER_GRPC_PORT", "8081"))

	c.ScheduleBranchHost = cast.ToString(getOrReturnDefault("SCHEDULE_SERVICE_HOST", "localhost"))
	c.ScheduleBranchPort = cast.ToString(getOrReturnDefault("SCHEDULE_GRPC_PORT", "8082"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	if os.Getenv(key) == "" {
		return defaultValue
	}

	return os.Getenv(key)
}
