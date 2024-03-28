package randPlace

import (
	"math/rand"

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
	if len(g.ValidSlots) == 0 {
		return -1, &ErrNoValidSlots{}
	}
	return g.ValidSlots[rand.Intn(len(g.ValidSlots))], nil
}
