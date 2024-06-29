package repositories

import (
	"fmt"

	"github.com/org/sample_go_worker/src/data"
	"github.com/org/sample_go_worker/src/data/models/dax"
)

type GenericItemsRepository interface {
	AddItem(name string) (*dax.Item, error)
	GetItem(id string) (*dax.Item, error)
	DeleteItem(id string) error
}

type ItemsRepository struct {
	Database data.Database
}

func NewItemsRepository(db data.Database) ItemsRepository {
	repo := ItemsRepository{
		Database: db,
	}
	return repo
}

func (r ItemsRepository) AddItem(name string) (*dax.Item, error) {

	item := dax.Item{DisplayName: name}

	result := r.Database.SqlDb.Create(&item)
	if result.Error != nil {
		fmt.Printf("cannot create new item in DB: %v", result.Error)
		return nil, result.Error
	}

	fmt.Println(result.RowsAffected)

	fmt.Println(item)

	return &item, nil

}

func (r ItemsRepository) GetItem(id string) (*dax.Item, error) {

	var item dax.Item

	err := r.Database.SqlDb.First(&item, "id = ?", id)
	if err.Error != nil {
		fmt.Printf("cannot create new item in DB: %v", err.Error)
		return nil, err.Error
	}

	return &item, nil

}

func (r ItemsRepository) DeleteItem(id string) error {

	err := r.Database.SqlDb.Delete(&dax.Item{}, id)
	if err.Error != nil {
		fmt.Printf("cannot create new item in DB: %v", err.Error)
		return err.Error
	}

	return nil

}
