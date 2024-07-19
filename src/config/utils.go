package config

import (
	"log/slog"
)

// Logs to console the current app configuration, as taken from the environment.
// Useful to ensure that the application has the correct configuration in the various hosting environments.
//
// Supports suppresing certain secret values through the SuppressSecretsInConfigLogs flag.
func LogAppConfiguration(logger slog.Logger) {
	if SuppressSecretsInConfigLogs == 1 || SuppressSecretsInConfigLogsErr != nil {
		logger.Info("[SUPPRESSED] Running worker with configuration:\n",
			// Runtime utils
			slog.String("ENVIRONMENT", Environment),
			slog.String("PORT", Port),
			// Connection strings (SUPPRESS VALUES)
			slog.String("MYSQL_CONNECTION_STRING", "***"),
		)
	} else {
		logger.Info("[PLAINTEXT] Running worker with configuration:\n",
			// Runtime utils
			slog.String("ENVIRONMENT", Environment),
			slog.String("PORT", Port),
			// Connection strings (USE ACTUAL VALUES)
			slog.String("MYSQL_CONNECTION_STRING", MySqlConnectionString),
		)
	}
}
