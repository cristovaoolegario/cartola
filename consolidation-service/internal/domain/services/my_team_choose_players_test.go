package services

import (
	"reflect"
	"testing"

	"github.com/cristovaoolegario/cartola/consolidation-service/internal/domain/entity"
)

func TestChoosePlayers(t *testing.T) {
	myTeam := entity.NewMyTeam("123", "test team")
	player1 := entity.NewPlayer("1", "Test player 1", 10.0)
	player2 := entity.NewPlayer("2", "Test player 2", 5.0)
	playersToBuy := []entity.Player{}
	myPlayers := []entity.Player{}

	t.Run("Should thrown an error When there's not enough founds to do the transactions", func(t *testing.T) {
		playersToBuy = []entity.Player{*player1}
		myPlayers = []entity.Player{*player2}
		myTeam.Players = append(myTeam.Players, player2.ID)

		err := ChoosePlayers(myTeam, myPlayers, playersToBuy)

		if err == nil {
			t.Error("Expected an error to be thrown")
		}

		if err.Error() != "not enough founds" {
			t.Errorf("Expected to thrown 'not enough founds' error, got %s", err.Error())
		}
	})

	t.Run("Should match all players IDs  and Score in the Players list When transactions are doable", func(t *testing.T) {
		playersToBuy = []entity.Player{*player2}
		myPlayers = []entity.Player{*player1}
		myTeam.Players = []string{}
		myTeam.Players = append(myTeam.Players, player1.ID)

		err := ChoosePlayers(myTeam, myPlayers, playersToBuy)

		expectedScore := player1.Price - player2.Price
		expectedPlayers := []string{"2"}

		if err != nil {
			t.Errorf("Not expected an error to be thrown, got %s ", err.Error())
		}

		if myTeam.Score != expectedScore {
			t.Errorf("Expected score to be %f, got %f", expectedScore, myTeam.Score)
		}

		if !reflect.DeepEqual(myTeam.Players, expectedPlayers) {
			t.Errorf("Expected players to be %s, got %s", expectedPlayers, myTeam.Players)
		}
	})

}
