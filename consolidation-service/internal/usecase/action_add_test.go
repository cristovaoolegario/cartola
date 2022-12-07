package usecase

import (
	"testing"

	"github.com/cristovaoolegario/cartola/consolidation-service/internal/domain/entity"
	"github.com/cristovaoolegario/cartola/consolidation-service/internal/infra/repository"
)

func TestAddAction(t *testing.T) {
	ctx, db := repository.SetupTestDb("../../")
	uow := repository.SetupTestUoW(ctx, db)
	err := repository.PopulateDb("../../", db)
	actionTable := entity.ActionTable{}
	actionTable.Init()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	t.Run("Execute ActionAdd UC", func(t *testing.T) {

		input := ActionAddInput{
			MatchID:  "1",
			TeamID:   "1",
			PlayerID: "1",
			Minute:   10,
			Action:   "goal",
		}

		uc := ActionAddUseCase{
			Uow:         uow,
			ActionTable: &actionTable,
		}

		t.Run("Should return an error When match doesn't exist", func(t *testing.T) {
			input.MatchID = "1000"

			err = uc.Execute(ctx, input)

			if err == nil {
				t.Error("Expected Match not to be registered")
			}
		})

		t.Run("Should return an error When action doesn't exist", func(t *testing.T) {
			input.MatchID = "1"
			input.Action = "not_an_action"

			err = uc.Execute(ctx, input)

			if err == nil || err.Error() != "action not found" {
				t.Error("Expected action not to be registered", err.Error())
			}
		})

		t.Run("Should return an error When player doesn't exist", func(t *testing.T) {
			input.MatchID = "1"
			input.Action = "goal"
			input.PlayerID = "13"

			err = uc.Execute(ctx, input)

			if err == nil || err.Error() != "sql: no rows in result set" {
				t.Error("Expected player not to be registered", err.Error())
			}
		})

		t.Run("Should return an error When MyTeam doesn't exist", func(t *testing.T) {
			input.MatchID = "1"
			input.Action = "goal"
			input.PlayerID = "1"
			input.TeamID = "13"

			err = uc.Execute(ctx, input)

			if err == nil || err.Error() != "sql: no rows in result set" {
				t.Error("Expected Myteam not to be registered", err.Error())
			}
		})

		t.Run("Should run successfully When all entities exist", func(t *testing.T) {
			input.MatchID = "1"
			input.Action = "goal"
			input.PlayerID = "1"
			input.TeamID = "1"

			err = uc.Execute(ctx, input)

			if err != nil {
				t.Error("Got an error:", err.Error())
			}
		})

	})
}
