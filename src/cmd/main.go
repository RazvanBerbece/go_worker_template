package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/org/sample_go_worker/src/config"
	"github.com/org/sample_go_worker/src/data"
	"github.com/org/sample_go_worker/src/data/repositories"
	"github.com/org/sample_go_worker/src/services"
)

// Loggers
var ConsoleLogger = slog.New(slog.NewTextHandler(os.Stdout, nil))

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

	// Sample app logic starts here
	_, err := ItemsService.CreateItem("SampleItem")
	if err != nil {
		fmt.Println(err)
	}
}
