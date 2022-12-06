package entity

import "testing"

func TestTeam(t *testing.T) {
	team := NewTeam("123", "Test team")
	player := NewPlayer("123", "Test player name", 10.0)
	team.AddPlayer(player)

	t.Run("Should add player to team When add player is called", func(t *testing.T) {
		if len(team.Players) < 1 {
			t.Errorf("Should've added a player to the team, any player was found in the team.")
		}
	})

	team.RemovePlayer(player)

	t.Run("Should remove player from team When remove player is called", func(t *testing.T) {
		if len(team.Players) > 0 {
			t.Errorf("Should'nt have any player in the team, found %d", len(team.Players))
		}
	})
}
