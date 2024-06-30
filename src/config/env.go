package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

// Runtime utils
var Environment = os.Getenv("ENVIRONMENT")
var LogAppConfigAtStartup, _ = strconv.Atoi(os.Getenv("LOG_APP_CONFIG_AT_STARTUP"))
var SuppressSecretsInConfigLogs, SuppressSecretsInConfigLogsErr = strconv.Atoi(os.Getenv("SUPRESS_SECRETS_IN_STARTUP_LOGS"))

// Connection strings for various external dependencies
var MySqlConnectionString = os.Getenv("MYSQL_CONNECTION_STRING")
