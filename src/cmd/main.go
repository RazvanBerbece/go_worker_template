package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/org/sample_go_worker/src/config"
	"github.com/org/sample_go_worker/src/data"
	"github.com/org/sample_go_worker/src/data/repositories"
	"github.com/org/sample_go_worker/src/services"
)

// Loggers
var ConsoleLogger = slog.Default()

// Infrastructure
var Database = data.NewDatabaseService(config.MySqlConnectionString, true, true)

// Services
var ItemsService = services.ItemsService{
	Logger:          *ConsoleLogger,
	ItemsRepository: repositories.NewItemsRepository(Database),
}

func main() {

	// If enabled, output to console the app configuration
	if config.LogAppConfigAtStartup == 1 {
		go config.LogAppConfiguration(*ConsoleLogger)
	}

	// e.g: Blocking call to handle requests
	// This could be replaced with different logic, but HTTP request handling is given as an example.
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.Port), nil); err != nil {
		ConsoleLogger.Error("Failed to start backend worker:", slog.String("error", err.Error()))
	}
}
