package repositories

import (
	"fmt"

	"github.com/org/sample_go_worker/src/data"
	"github.com/org/sample_go_worker/src/data/models/dax"
)

type GenericItemsRepository interface {
	AddItem(name string) (*dax.Item, error)
	GetItem(id uint) (*dax.Item, error)
	DeleteItem(id uint) error
}

type ItemsRepository struct {
	Database data.Database
}

func NewItemsRepository(db data.Database) GenericItemsRepository {
	repo := ItemsRepository{Database: db}
	return repo
}

func (r ItemsRepository) AddItem(name string) (*dax.Item, error) {

	item := dax.Item{DisplayName: name}

	result := r.Database.SqlDb.Create(&item)
	if result.Error != nil {
		fmt.Printf("cannot create new item in DB: %v\n", result.Error)
		return nil, result.Error
	}

	return &item, nil

}

func (r ItemsRepository) GetItem(id uint) (*dax.Item, error) {

	var item dax.Item

	err := r.Database.SqlDb.First(&item, "id = ?", id)
	if err.Error != nil {
		fmt.Printf("cannot retrieve item with ID %d from DB: %v\n", id, err.Error)
		return nil, err.Error
	}

	return &item, nil

}

func (r ItemsRepository) DeleteItem(id uint) error {

	err := r.Database.SqlDb.Delete(&dax.Item{}, id)
	if err.Error != nil {
		fmt.Printf("cannot delete item from DB: %v\n", err.Error)
		return err.Error
	}

	return nil

}
