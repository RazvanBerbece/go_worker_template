package main

import (
	"fmt"

	"github.com/org/sample_go_worker/src/config"
	"github.com/org/sample_go_worker/src/data"
	"github.com/org/sample_go_worker/src/data/repositories"
)

// Infrastructure
// var Logger =
var Database = data.NewDatabaseService(config.MySqlConnectionString, true)

// Repositories
var ItemsRepository = repositories.NewItemsRepository(Database)

// Services
// var ItemsService = services.NewItemsService(Logger, ItemsRepository)

func main() {
	fmt.Printf("Running in environment -> %s\n", config.Environment)
	fmt.Printf("	with DB connection string -> %s\n", config.MySqlConnectionString)

	_, err := ItemsRepository.AddItem("Test Item")
	if err != nil {
		fmt.Println(err)
	}
}
