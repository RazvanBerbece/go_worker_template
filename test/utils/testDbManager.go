package test_utils

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/org/sample_go_worker/src/data"
	"github.com/org/sample_go_worker/src/data/models/dax"
)

// Wrapper around a data.Database type which provides util extensions for interacting with the DB in the automated testing context.
// Usually, the contained database instance will be configured with the CI environment connection strings, flags, etc.
type TestDatabaseManager struct {
	Database data.Database
}

func (m TestDatabaseManager) AddItemToDb() error {

	item := dax.Item{DisplayName: gofakeit.Name()}

	result := m.Database.SqlDb.Create(&item)
	if result.Error != nil {
		fmt.Printf("cannot create new item in DB: %v\n", result.Error)
		return result.Error
	}

	return nil

}

func (m TestDatabaseManager) AddItemWithNameToDb(name string) error {

	item := dax.Item{DisplayName: name}

	result := m.Database.SqlDb.Create(&item)
	if result.Error != nil {
		fmt.Printf("cannot create new item in DB: %v\n", result.Error)
		return result.Error
	}

	return nil

}
