package usecase

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/cristovaoolegario/cartola/consolidation-service/internal/infra/repository"
	"github.com/cristovaoolegario/cartola/consolidation-service/pkg/uow"
)

func TestPlayer(t *testing.T) {
	ctx := context.Background()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Errorf("An error has occur when connecting to the database: %s", err.Error())
	}

	err = repository.RunDbInit("../../", db)

	if err != nil {
		t.Errorf("An error has occur when running the Migrations: %s", err.Error())
	}

	defer db.Close()

	uow, err := uow.NewUow(ctx, db)
	if err != nil {
		t.Errorf("An error has occur when setting up Unit of Work: %s", err.Error())
	}
	repository.RegisterRepositories(uow)

	t.Run("Execute AddPlayers UC", func(t *testing.T) {
		ucInput := AddPlayerInput{
			ID:           "1",
			Name:         "AddPlayerInput name",
			InitialPrice: 10.0,
		}

		uc := AddPlayerUseCase{
			Uow: uow,
		}

		err = uc.Execute(ctx, ucInput)

		if err != nil {
			t.Errorf("Not expected an error to happened, got %s", err.Error())
		}

		err = uc.Execute(ctx, ucInput)

		if err == nil {
			t.Error("Expected 'UNIQUE constraint failed: players.id' error to thrown, got none")
		}

	})
}
