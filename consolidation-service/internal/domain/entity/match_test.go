package entity

import (
	"fmt"
	"math/rand"
	"testing"
)

func FuzzClient_MatchResult_GetResult(f *testing.F) {
	seeding := []int{2, 4, 6, 8, 10}
	r := rand.New(rand.NewSource(99))
	for _, seed := range seeding {
		f.Add(seed, r.Int())
	}

	f.Fuzz(func(t *testing.T, a int, b int) {
		t.Run("Should return the correct match result When GetResult is called", func(t *testing.T) {
			mr := NewMatchResult(a, b)

			result := mr.GetResult()
			expected := fmt.Sprintf("%d-%d", a, b)
			if result != expected {
				t.Errorf("Should've returned the proper match result %s, got %s", expected, result)
			}
		})
	})
}
