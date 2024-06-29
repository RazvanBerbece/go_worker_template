package config

import (
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

// Runtime utils
var Environment = os.Getenv("ENVIRONMENT")
var SuppressSecretsInConfigLogs, SuppressSecretsInConfigLogsErr = strconv.ParseBool(os.Getenv("SUPRESS_SECRETS_IN_CONFIG_LOGS"))

// Connection strings for various external dependencies
var MySqlConnectionString = os.Getenv("MYSQL_CONNECTION_STRING")

func LogAppConfiguration(logger slog.Logger) {

	if SuppressSecretsInConfigLogsErr != nil {
		SuppressSecretsInConfigLogs = false
	}

	// NON-SECRET APP CONFIGURATIONS
	// Runtime utils
	log.Printf("ENVIRONMENT=%s", Environment)

	// SECRET APP CONFIGURATIONS
	if SuppressSecretsInConfigLogs {
		// Connection strings
		log.Printf("MYSQL_CONNECTION_STRING=%s", "***")
	} else {
		// Connection strings
		log.Printf("MYSQL_CONNECTION_STRING=%s", MySqlConnectionString)
	}
}
