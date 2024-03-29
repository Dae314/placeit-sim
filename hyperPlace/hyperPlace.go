package hyperPlace

import (
	"github.com/Dae314/placeit-sim/game"
	"github.com/Dae314/placeit-sim/utils"
	"gonum.org/v1/gonum/stat/combin"
)

type ErrNoValidSlots struct{}

func (e *ErrNoValidSlots) Error() string {
	return "Cannot suggest placement if there are no valid slots."
}

type gap struct {
	min int
	max int
	len int
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

	var winRates []float64

	for _, v := range g.ValidSlots {
		testSlots := make([]int, len(g.Slots))
		_ = copy(testSlots, g.Slots)
		testSlots[v] = g.CurDraw
		winRates = append(winRates, calcWin(testSlots, len(g.Deck), countEmpty(testSlots)))
	}

	maxIdx := utils.MaxSlicei(winRates)

	return g.ValidSlots[maxIdx], nil
}

func calcWin(slots []int, deckSize int, draws int) float64 {
	gaps := getGaps(slots)
	win := float64(1)
	for _, g := range gaps {
		gapRange := g.max - g.min
		win = (float64(combin.Binomial(gapRange, g.len)) * float64(combin.Binomial(deckSize-gapRange, draws-g.len))) / float64(combin.Binomial(deckSize, draws))
	}
	return win
}

func getGaps(slots []int) []gap {
	var gaps []gap
	for i, v := range slots {
		gapStart := -1
		gapLen := 0
		if i == 0 && v == -1 {
			gapStart = 1
			gapLen++
		} else if i == len(slots)-1 && gapStart != -1 && v == -1 {
			gapLen++
			gaps = append(gaps, gap{min: gapStart, max: 1000, len: gapLen})
		} else if v != -1 && gapStart != -1 {
			gaps = append(gaps, gap{min: gapStart, max: v, len: gapLen})
			gapStart = -1
		} else if v == -1 && gapStart == -1 {
			gapStart = slots[i-1]
			gapLen++
		} else if v == -1 && gapStart != -1 {
			gapLen++
		}
	}
	return gaps
}

func countEmpty(slots []int) int {
	i := 0
	for _, v := range slots {
		if v == -1 {
			i++
		}
	}
	return i
}
