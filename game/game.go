package game

type PlaceItGame struct {
	slots []int
}

func NewGame() PlaceItGame {
	return PlaceItGame{slots: make([]int, 20)}
}
