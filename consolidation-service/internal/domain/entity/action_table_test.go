package entity

import "testing"

func TestActionTable(t *testing.T) {
	at := ActionTable{}
	at.Init()

	t.Run("Should return action value When it exists", func(t *testing.T) {
		ac, err := at.GetScore("goal")

		if ac == 0 && err != nil {
			t.Errorf("Didn't expected an error to be thrown, got %s", err.Error())
		}
	})

	t.Run("Should return an error When action doesn't exist", func(t *testing.T) {
		ac, err := at.GetScore("not_an_action")

		if ac != 0 && err == nil {
			t.Errorf("Expected 'invalid action' error to happend, didn't got one")
		}
	})
}
