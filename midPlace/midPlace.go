package midPlace

import (
	"github.com/Dae314/placeit-sim/game"
)

type ErrNoValidSlots struct{}

func (e *ErrNoValidSlots) Error() string {
	return "Cannot suggest placement if there are no valid slots."
}

func GetPlacement(g *game.PlaceItGame) (int, error) {
	if g.State != game.PlaceState {
		return -1, &game.ErrGameInvalidState{
			State:  g.State,
			Action: "GetPlacement",
		}
	}
	switch len(g.ValidSlots) {
	case 0:
		return -1, &ErrNoValidSlots{}
	case 1:
		return g.ValidSlots[0], nil
	default:
		mid := len(g.ValidSlots) / 2
		return g.ValidSlots[mid-1], nil
	}
}
