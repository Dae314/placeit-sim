package hyperPlace

import (
	"github.com/Dae314/placeit-sim/game"
	"github.com/Dae314/placeit-sim/utils"
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
		winRates = append(winRates, calcWin(testSlots, g))
	}

	maxIdx := utils.MaxSlicei(winRates)

	return g.ValidSlots[maxIdx], nil
}

func calcWin(slots []int, testGame *game.PlaceItGame) float64 {
	gaps := getGaps(slots)
	deckSize := len(testGame.Deck)
	draws := countEmpty(testGame.Slots)
	gapCombinAccu := float64(1)
	gapRangeAccu := 0
	gapLenAccu := 0
	for _, g := range gaps {
		gapRange := calcGapRange(g, testGame.Deck)
		if gapRange < g.len {
			return 0
		}
		gapCombinAccu = gapCombinAccu * Choose(float64(gapRange), float64(g.len))
		gapRangeAccu = gapRangeAccu + gapRange
		gapLenAccu = gapLenAccu + g.len
	}
	allCombin := Choose(float64(deckSize), float64(draws))

	return gapCombinAccu / allCombin
}

func getGaps(slots []int) []gap {
	var gaps []gap
	gapStart := -1
	gapLen := 0
	for i, v := range slots {
		if i == 0 && v == -1 {
			gapStart = 0
			gapLen++
		} else if i == len(slots)-1 && gapStart != -1 && v == -1 {
			gapLen++
			gaps = append(gaps, gap{min: gapStart, max: 1001, len: gapLen})
		} else if v != -1 && gapStart != -1 {
			gaps = append(gaps, gap{min: gapStart, max: v, len: gapLen})
			gapStart = -1
			gapLen = 0
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

func calcGapRange(g gap, deck []int) int {
	i := 0
	for _, v := range deck {
		if v > g.min && v < g.max {
			i++
		}
	}
	return i
}

func Choose(n, k float64) float64 {
	if k > n/2 {
		k = n - k
	}
	b := float64(1)
	for i := float64(1); i <= k; i++ {
		b = (n - k + i) * b / i
	}
	return b
}
