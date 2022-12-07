package repository

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/cristovaoolegario/cartola/consolidation-service/internal/domain/entity"
)

func TestPlayerRepository(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Errorf("An error has occur when connecting to the database: %s", err.Error())
	}

	err = RunDbInit("../../../", db)

	if err != nil {
		t.Errorf("An error has occur when running the Migrations: %s", err.Error())
	}

	defer db.Close()

	t.Run("Create", func(t *testing.T) {
		repo := NewPlayerRepository(db)

		err := repo.Create(context.TODO(), entity.NewPlayer("1", "Test player", 10.0))

		if err != nil {
			t.Errorf("An error has occur when creating a new Player: %s", err.Error())
		}
	})

	t.Run("FindByID", func(t *testing.T) {
		t.Run("Should return Player When a Player exists with ID", func(t *testing.T) {
			repo := NewPlayerRepository(db)

			player, err := repo.FindByID(context.TODO(), "1")

			if player == nil && err != nil {
				t.Errorf("An error has occur when searching for an existing player: %s", err.Error())
			}

		})
		t.Run("Should return Error When a Player don't exists with ID", func(t *testing.T) {
			repo := NewPlayerRepository(db)

			player, err := repo.FindByID(context.TODO(), "2")

			if player != nil && err == nil {
				t.Error("Should've returned an error, find a player")
			}
		})
	})
}
