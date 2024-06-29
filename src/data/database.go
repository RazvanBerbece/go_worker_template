package data

import (
	"fmt"

	"github.com/org/sample_go_worker/src/data/models/dax"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	ConnectionString string
	SqlDb            *gorm.DB

	// Configurations
	RetryCount     int
	AutoMigrations bool
}

// Returns a data service connected to the database passed in through the connection string.
// Runs migrations based on the GORM schema if the flag is set to true.
func NewDatabaseService(connectionString string, retryCount int, autoMigrations bool) Database {

	var database = Database{
		ConnectionString: connectionString,
		AutoMigrations:   autoMigrations,
		RetryCount:       retryCount,
	}

	err := database.Connect()
	if err != nil {
		panic(err)
	}

	return database

}

func (dbm *Database) Connect() error {

	var db *gorm.DB

	// Connect with a basic retry strategy
	// Attempt N times to connect
	for range dbm.RetryCount {
		var err error
		db, err = gorm.Open(mysql.Open(dbm.ConnectionString), &gorm.Config{})
		if err == nil {
			break
		} else {
			fmt.Printf("an error ocurred while connecting to the DB: %v\n", err)
			fmt.Println("retrying to establish connection to the DB...")
		}
	}

	if dbm.AutoMigrations {
		db.AutoMigrate(&dax.Item{})
	}

	dbm.SqlDb = db

	return nil

}
