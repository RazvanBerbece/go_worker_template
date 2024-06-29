package config

import (
	"log"
	"log/slog"
)

func LogAppConfiguration(logger slog.Logger) {

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
