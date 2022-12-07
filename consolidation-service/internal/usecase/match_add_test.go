package usecase

import (
	"testing"
	"time"

	"github.com/cristovaoolegario/cartola/consolidation-service/internal/infra/repository"
)

func TestMatchAddUseCase(t *testing.T) {
	ctx, db := repository.SetupTestDb("../../")
	uow := repository.SetupTestUoW(ctx, db)
	err := repository.PopulateDb("../../", db)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	t.Run("Execute MatchAdd UC with sucess", func(t *testing.T) {
		input := MatchInput{
			ID:      "3",
			Date:    time.Now(),
			TeamAID: "1",
			TeamBID: "2",
		}

		uc := MatchUseCase{
			Uow: uow,
		}

		err = uc.Execute(ctx, input)

		if err != nil {
			t.Errorf("Not expected an error to happened, got %s", err.Error())
		}

		t.Run("Should thown an error When match is already registered", func(t *testing.T) {
			err = uc.Execute(ctx, input)

			if err == nil || err.Error() != "UNIQUE constraint failed: matches.id" {
				t.Errorf("Expected 'UNIQUE constraint failed: matches.id' error to thrown, got none")
			}
		})

		t.Run("Should return an error When Team A or Team B doesn't exist", func(t *testing.T) {
			input.TeamAID = "12"
			err = uc.Execute(ctx, input)

			if err == nil {
				t.Error("Expected Team A not to be registered")
			}

			input.TeamAID = "1"
			input.TeamBID = "15"
			err = uc.Execute(ctx, input)

			if err == nil {
				t.Error("Expected Team B not to be registered")
			}
		})
	})
}
