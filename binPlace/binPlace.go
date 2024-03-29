package binPlace

import (
	"math"

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

	minIdx := g.ValidSlots[0]
	maxIdx := g.ValidSlots[validLen-1]

	var min int
	var max int

	if minIdx == 0 {
		min = 0
	} else {
		min = g.Slots[minIdx-1]
	}

	if maxIdx == len(g.Slots)-1 {
		max = 1001
	} else {
		max = g.Slots[maxIdx+1]
	}

	binRange := max - min
	binStep := int(math.Ceil(float64(binRange) / float64(validLen)))
	i := int(math.Floor(float64(g.CurDraw-min) / float64(binStep)))

	return g.ValidSlots[i], nil
}
