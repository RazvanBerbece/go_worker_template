package unit

import (
	"log/slog"
	"testing"

	"github.com/org/sample_go_worker/src/services"
	mocks "github.com/org/sample_go_worker/test/integration/service/mocks/repos"
)

func TestExample(t *testing.T) {

	// Arrange
	var _ = services.ItemsService{
		Logger:          slog.Logger{},
		ItemsRepository: mocks.NewMockItemsRepository(),
	}

}
