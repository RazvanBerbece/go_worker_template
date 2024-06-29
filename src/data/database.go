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
	AutoMigrations bool
}

// Returns a data service connected to the database passed in through the connection string.
// Runs migrations based on the GORM schema if the flag is set to true.
func NewDatabaseService(connectionString string, autoMigrations bool) Database {

	var database = Database{
		ConnectionString: connectionString,
		AutoMigrations:   autoMigrations,
	}

	database.Connect()

	return database

}

func (dbm *Database) Connect() error {

	db, err := gorm.Open(mysql.Open(dbm.ConnectionString), &gorm.Config{})
	if err != nil {
		fmt.Printf("an error ocurred while connecting to the DB: %v\n", err)
		return err
	}

	if dbm.AutoMigrations {
		db.AutoMigrate(&dax.Item{})
	}

	dbm.SqlDb = db

	return nil

}
