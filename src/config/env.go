package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

//
// New environment variables (app configuration key-value pairs) should be added here.
//
// Once an environment variable is added in here,
// it could also then be integrated in the LogAppConfiguration() function in utils.go.
//

var _ = godotenv.Load(".env")

// Runtime utils
var Environment = os.Getenv("ENVIRONMENT")
var Port = os.Getenv("PORT")

// Runtime flags
var LogAppConfigAtStartup, _ = strconv.Atoi(os.Getenv("LOG_APP_CONFIG_AT_STARTUP"))
var SuppressSecretsInConfigLogs, SuppressSecretsInConfigLogsErr = strconv.Atoi(os.Getenv("SUPRESS_SECRETS_IN_STARTUP_LOGS"))

// Connection strings for various external dependencies
var MySqlConnectionString = os.Getenv("MYSQL_CONNECTION_STRING")
