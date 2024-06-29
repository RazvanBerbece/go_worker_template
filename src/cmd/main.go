package main

import (
	"fmt"

	"github.com/org/sample_go_worker/src/config"
)

func main() {
	fmt.Printf("Running in environment -> %s\n", config.Environment)
	fmt.Printf("	with DB connection string -> %s\n", config.MySqlConnectionString)
}
