package hyperPlace

import (
	"testing"

	"github.com/Dae314/placeit-sim/game"
	"github.com/Dae314/placeit-sim/utils"
)

func TestBinPlacement(t *testing.T) {
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
