package game

import (
	"fmt"
	"math/rand"
)

const deckSize = 1000
const slotSize = 20

type GameState int

const (
	DrawState GameState = iota
	PlaceState
	WinState
	LoseState
)

type ErrGameInvalidState struct {
	state  GameState
	action string
}

func (e *ErrGameInvalidState) Error() string {
	stateStrings := []string{"DrawState", "PlaceState", "WinState", "LoseState"}
	return fmt.Sprintf("State %s is not a valid state for %s", stateStrings[e.state], e.action)
}

type ErrPlaceOutofBounds struct {
	index int
}

func (e *ErrPlaceOutofBounds) Error() string {
	return fmt.Sprintf("Attempted to place at index %d, which is not a valid placement slot.", e.index)
}

type PlaceItGame struct {
	Slots      []int
	State      GameState
	CurDraw    int
	ValidSlots []int
	Deck       []int
}

func NewGame() PlaceItGame {
	deck := make([]int, deckSize)
	n := 1
	for i := range deck {
		deck[i] = n
		n++
	}
	slots := make([]int, slotSize)
	for i := range slots {
		slots[i] = -1
	}
	return PlaceItGame{
		Slots:   slots,
		State:   DrawState,
		CurDraw: -1,
		Deck:    deck,
	}
}

func (g *PlaceItGame) Draw() error {
	if g.State != DrawState {
		return &ErrGameInvalidState{state: g.State, action: "Draw"}
	}
	randIdx := rand.Intn(len(g.Deck))
	g.CurDraw = g.Deck[randIdx]
	g.Deck = removeFromSlice(g.Deck, randIdx)
	g.ValidSlots = validSlots(g)
	g.State = checkGameover(g)
	return nil
}

func (g *PlaceItGame) Place(i int) error {
	if g.State != PlaceState {
		return &ErrGameInvalidState{state: g.State, action: "Place"}
	}
	if !contains(g.ValidSlots, i) {
		return &ErrPlaceOutofBounds{index: i}
	}
	if g.Slots[i] != -1 {
		return &ErrPlaceOutofBounds{index: i}
	}
	g.Slots[i] = g.CurDraw
	g.CurDraw = -1
	g.State = DrawState
	return nil
}

func validSlots(g *PlaceItGame) []int {
	type pair struct {
		index int
		value int
	}
	var validSlots []int
	var vals []pair
	empty := true
	for i, v := range g.Slots {
		if v > 0 {
			vals = append(vals, pair{i, v})
			empty = false
		}
	}
	if empty {
		return []int{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		}
	}
	leftMax := -1
	leftIndex := -1
	rightMin := -1
	rightIndex := -1
	for _, v := range vals {
		if v.value < g.CurDraw && (leftMax == -1 || v.value > leftMax) {
			leftMax = v.value
			leftIndex = v.index
		}
		if v.value > g.CurDraw && (rightMin == -1 || v.value < rightMin) {
			rightMin = v.value
			rightIndex = v.index
		}
	}

	if leftMax == -1 && rightMin != -1 {
		for i := 0; i < rightIndex; i++ {
			if g.Slots[i] == -1 {
				validSlots = append(validSlots, i)
			}
		}
	} else if leftMax != -1 && rightMin == -1 {
		for i := leftIndex; i < len(g.Slots); i++ {
			if g.Slots[i] == -1 {
				validSlots = append(validSlots, i)
			}
		}
	} else {
		for i := leftIndex; i < rightIndex; i++ {
			if g.Slots[i] == -1 {
				validSlots = append(validSlots, i)
			}
		}
	}
	return validSlots
}

func checkGameover(g *PlaceItGame) GameState {
	if len(g.ValidSlots) != 0 {
		return PlaceState
	}
	for _, v := range g.Slots {
		if v == -1 {
			return LoseState
		}
	}
	return WinState
}

func removeFromSlice(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func contains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
