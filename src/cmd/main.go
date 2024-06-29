package main

import (
	"fmt"

	"github.com/org/go_worker/src/config"
)

func main() {
	fmt.Printf("Running in < %s >\n", config.Environment)
	fmt.Printf("With DB connection string: < %s >\n", config.MySqlConnectionString)
}
