package randPlace

import (
	"testing"

	"github.com/Dae314/placeit-sim/game"
)

func TestRandomPlacement(t *testing.T) {
	testGame := game.NewGame()
	testGame.Draw()
	myDraw := testGame.CurDraw
	idx, err := GetPlacement(&testGame)
	testGame.Place(idx)
	if err != nil {
		t.Errorf("Got %v when error wasn't expected", err)
	}
	checkContains(t, testGame.Slots, myDraw)
}

func checkContains[T comparable](t testing.TB, got []T, want T) {
	t.Helper()
	if !contains(got, want) {
		t.Errorf("%v does not contain %v", got, want)
	}
}

func contains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
