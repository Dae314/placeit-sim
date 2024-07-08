package simpleBinPlace

import (
	"errors"
	"testing"

	"github.com/Dae314/placeit-sim/game"
	"github.com/Dae314/placeit-sim/utils"
)

func TestSimpleBinPlacement(t *testing.T) {
	testGame := game.NewGame()
	testGame.Draw()
	myDraw := testGame.CurDraw
	idx, err := GetPlacement(&testGame)
	testGame.Place(idx)
	if err != nil {
		t.Errorf("Got %v when error wasn't expected", err)
	}
	utils.CheckContains(t, testGame.Slots, myDraw)
}

func TestSimpleBinPlacementError(t *testing.T) {
	testGame := game.NewGame()
	testGame.Slots = []int{-1, 60, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	testGame.State = game.PlaceState
	testGame.CurDraw = 55
	testGame.ValidSlots = []int{0}
	_, err := GetPlacement(&testGame)
	if err == nil {
		t.Errorf("Did not get error when error was expected")
	} else {
		expected := &ErrEndGame{}
		if !errors.As(err, &expected) {
			t.Errorf("Expected to get %v error, instead received: %v", expected, err)
		}
	}
}
