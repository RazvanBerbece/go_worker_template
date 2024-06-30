package unit

import (
	"os"
	"testing"

	"github.com/org/sample_go_worker/src/data"
	"github.com/org/sample_go_worker/src/data/models/dax"
	"github.com/org/sample_go_worker/src/data/repositories"
	test_utils "github.com/org/sample_go_worker/test/utils"
	"gorm.io/gorm"
)

var TestDatabaseManager = test_utils.TestDatabaseManager{
	Database: data.NewDatabaseService(os.Getenv("MYSQL_CONNECTION_STRING"), true, true),
}

func TestExample(t *testing.T) {

	var sut = repositories.NewItemsRepository(TestDatabaseManager.Database)

	// Arrange
	var itemName = "TestItem"
	var itemName2 = "TestItem2"
	TestDatabaseManager.AddItemWithNameToDb(itemName)
	TestDatabaseManager.AddItemWithNameToDb(itemName2)

	cases := []struct {
		input          uint
		expectedOutput *dax.Item
	}{
		{1, &dax.Item{Model: gorm.Model{ID: 1}, DisplayName: itemName}},
		{2, &dax.Item{Model: gorm.Model{ID: 2}, DisplayName: itemName2}},
		{99, nil},
	}

	for _, c := range cases {

		// Act
		output, _ := sut.GetItem(c.input)

		// Assert
		if c.expectedOutput == nil {
			if output != nil {
				t.Errorf("incorrect actual output for ID %d: expected null", c.input)
			}
		} else {
			if output.ID != c.expectedOutput.ID {
				t.Errorf("incorrect actual output for ID %d: expected ID %d, got %d", c.input, c.expectedOutput.ID, output.ID)
			}
			if output.DisplayName != c.expectedOutput.DisplayName {
				t.Errorf("incorrect actual output for ID %d: expected DisplayName %s, got %s", c.input, c.expectedOutput.DisplayName, output.DisplayName)
			}
		}
	}

}
