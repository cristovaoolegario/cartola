package usecase

import (
	"testing"

	"github.com/cristovaoolegario/cartola/consolidation-service/internal/infra/repository"
)

func TestAddMyTeamUseCase(t *testing.T) {
	ctx, db := repository.SetupTestDb("../../")
	uow := repository.SetupTestUoW(ctx, db)

	defer db.Close()

	t.Run("Execute AddMyTeam UC", func(t *testing.T) {
		addMyTeamInput := AddMyTeamInput{
			ID:    "1",
			Name:  "AddMyTeamInput test",
			Score: 0,
		}

		uc := AddMyTeamUseCase{
			Uow: uow,
		}

		err := uc.Execute(ctx, addMyTeamInput)

		if err != nil {
			t.Errorf("Not expected an error to happened, got %s", err.Error())
		}

		err = uc.Execute(ctx, addMyTeamInput)

		if err == nil {
			t.Errorf("Expected 'UNIQUE constraint failed: my_team.id' error to thrown, got none")
		}
	})
}
