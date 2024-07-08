package simpleBinRand

import (
	"math"
	"math/rand"
	"slices"

	"github.com/Dae314/placeit-sim/game"
)

type ErrNoValidSlots struct{}

func (e *ErrNoValidSlots) Error() string {
	return "Cannot suggest placement if there are no valid slots."
}

func GetPlacement(g *game.PlaceItGame) (int, error) {
	validLen := len(g.ValidSlots)
	if g.State != game.PlaceState {
		return -1, &game.ErrGameInvalidState{
			State:  g.State,
			Action: "GetPlacement",
		}
	}
	if validLen == 0 {
		return -1, &ErrNoValidSlots{}
	}
	if validLen == 1 {
		return g.ValidSlots[0], nil
	}

	s := int(math.Ceil(float64(g.CurDraw)/(float64(len(g.Deck))/float64(len(g.Slots))))) - 1
	if s > len(g.Slots)-1 {
		s = len(g.Slots) - 1
	}

	if !slices.Contains(g.ValidSlots, s) {
		return g.ValidSlots[rand.Intn(validLen)], nil
	}

	return s, nil
}
