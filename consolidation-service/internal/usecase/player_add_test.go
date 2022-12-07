package usecase

import (
	"testing"

	"github.com/cristovaoolegario/cartola/consolidation-service/internal/infra/repository"
)

func TestPlayer(t *testing.T) {
	ctx, db := repository.SetupTestDb("../../")
	uow := repository.SetupTestUoW(ctx, db)

	defer db.Close()

	t.Run("Execute AddPlayers UC", func(t *testing.T) {
		ucInput := AddPlayerInput{
			ID:           "1",
			Name:         "AddPlayerInput name",
			InitialPrice: 10.0,
		}

		uc := AddPlayerUseCase{
			Uow: uow,
		}

		err := uc.Execute(ctx, ucInput)

		if err != nil {
			t.Errorf("Not expected an error to happened, got %s", err.Error())
		}

		err = uc.Execute(ctx, ucInput)

		if err == nil {
			t.Error("Expected 'UNIQUE constraint failed: players.id' error to thrown, got none")
		}

	})
}
