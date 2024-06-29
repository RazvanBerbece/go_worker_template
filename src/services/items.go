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
		fmt.Println(err)
		return nil, err
	}

	return item, nil

}
