package repositories

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/org/sample_go_worker/src/data"
	"github.com/org/sample_go_worker/src/data/models/dax"
)

type GenericItemsRepository interface {
	AddItem(name string) (int64, error)
	GetItem(id string) (*dax.Item, error)
	DeleteItem(id string) (int64, error)
}

type ItemsRepository struct {
	DbContext data.GenericDatabaseContext
}

func NewItemsRepository(connString string) ItemsRepository {
	repo := ItemsRepository{&data.DatabaseContext{
		ConnectionString: connString,
	}}
	repo.DbContext.Connect()
	return repo
}

func (r ItemsRepository) AddItem(name string) (int64, error) {

	stockItem := &dax.Item{
		Id:          uuid.New().String(),
		DisplayName: name,
	}

	stmt, err := r.DbContext.GetConfiguredSqlDb().Prepare(`
		INSERT INTO 
			Items(
				id, 
				displayName
			)
		VALUES(?, ?);`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(stockItem.Id, stockItem.DisplayName)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil

}

func (r ItemsRepository) GetItem(id string) (*dax.Item, error) {

	query := "SELECT * FROM Items WHERE id = ?"
	row := r.DbContext.GetConfiguredSqlDb().QueryRow(query, id)

	var item dax.Item
	err := row.Scan(&item.Id,
		&item.DisplayName)

	if err != nil {
		return nil, fmt.Errorf("an error ocurred while retrieving item with ID `%s`", id)
	}

	return &item, nil

}

func (r ItemsRepository) DeleteItem(id string) (int64, error) {

	stmt, err := r.DbContext.GetConfiguredSqlDb().Prepare(`DELETE FROM Items WHERE id = ?`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
