package config

import (
	"os"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

// Runtime utils
var Environment = os.Getenv("ENVIRONMENT")

// Connection strings for various external dependencies
var MySqlConnectionString = os.Getenv("MYSQL_CONNECTION_STRING")
