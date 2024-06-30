package config

import (
	"log"
	"log/slog"
)

// Logs to console the current app configuration, as taken from the environment.
// Useful to ensure that the application has the correct configuration in the various hosting environments.
func LogAppConfiguration(logger slog.Logger) {

	// NON-SECRET APP CONFIGURATIONS
	// Runtime utils
	log.Printf("ENVIRONMENT=%s", Environment)

	// SECRET APP CONFIGURATIONS - Suppresed by default
	if SuppressSecretsInConfigLogs == 1 || SuppressSecretsInConfigLogsErr != nil {
		// Connection strings
		log.Printf("MYSQL_CONNECTION_STRING=%s", "***")
	} else {
		// Connection strings
		log.Printf("MYSQL_CONNECTION_STRING=%s", MySqlConnectionString)
	}
}
