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
var Database = data.NewDatabaseService(config.MySqlConnectionString, 5, true)

// Services
var ItemsService = services.ItemsService{
	Logger:          *ConsoleLogger,
	ItemsRepository: repositories.NewItemsRepository(Database),
}

func main() {

	fmt.Printf("Running in environment -> %s\n", config.Environment)
	fmt.Printf("	with DB connection string -> %s\n", config.MySqlConnectionString)

	_, err := ItemsService.CreateItem("SampleItem")
	if err != nil {
		fmt.Println(err)
	}
}
