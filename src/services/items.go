//
// This is an example of a service that depends on a logger and a repository.
// Other services can extend this model.
//

package services

import (
	"fmt"

	"github.com/org/sample_go_worker/src/data/models/dax"
	"github.com/org/sample_go_worker/src/data/repositories"

	"log/slog"
)

type ItemsService struct {
	Logger          slog.Logger
	ItemsRepository repositories.GenericItemsRepository
}

func (s ItemsService) CreateItem(name string) (*dax.Item, error) {

	// log := fmt.Sprintf("Creating item: %s", name)
	// s.Logger.Info(log)

	item, err := s.ItemsRepository.AddItem(name)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, fmt.Errorf("cannot create item with name %s", name)
	}

	return item, nil

}

func (s ItemsService) GetItem(id uint) (*dax.Item, error) {

	// log := fmt.Sprintf("Creating item: %s", name)
	// s.Logger.Info(log)

	item, err := s.ItemsRepository.GetItem(id)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, fmt.Errorf("cannot retrieve item with id %d", id)
	}

	return item, nil

}

func (s ItemsService) DeleteItem(id uint) error {

	// log := fmt.Sprintf("Creating item: %s", name)
	// s.Logger.Info(log)

	err := s.ItemsRepository.DeleteItem(id)
	if err != nil {
		s.Logger.Error(err.Error())
		return fmt.Errorf("cannot delete item with id %d", id)
	}

	return nil

}
