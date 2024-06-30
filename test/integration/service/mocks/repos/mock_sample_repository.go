package mocks

import (
	"github.com/org/sample_go_worker/src/data/models/dax"
	"github.com/org/sample_go_worker/src/data/repositories"
	"gorm.io/gorm"
)

type MockItemsRepository struct{}

func NewMockItemsRepository() repositories.GenericItemsRepository {
	repo := MockItemsRepository{}
	return repo
}

func (r MockItemsRepository) AddItem(name string) (*dax.Item, error) {

	item := dax.Item{DisplayName: name}

	return &item, nil

}

func (r MockItemsRepository) GetItem(id uint) (*dax.Item, error) {

	return &dax.Item{
		Model:       gorm.Model{ID: id},
		DisplayName: "TestName",
	}, nil

}

func (r MockItemsRepository) DeleteItem(id uint) error {

	return nil

}
