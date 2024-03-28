package game

import (
	"reflect"
	"testing"
)

func TestNewGame(t *testing.T) {
	game := NewGame()
	checkType(t, game, "game.PlaceItGame")
}

func TestDraw(t *testing.T) {
	game := NewGame()
	t.Run("Draw in DrawState", func(t *testing.T) {
		err := game.Draw()

		checkEquals(t, err, nil)
		checkEquals(t, game.State, PlaceState)
		checkNotNil(t, game.CurDraw)
	})
	t.Run("Draw in PlaceState", func(t *testing.T) {
		err := game.Draw()

		checkType(t, err, "*game.ErrGameInvalidState")
	})
}

func TestPlace(t *testing.T) {
	slots := []int{
		-1, -1, 120, -1, -1, 200, -1, -1, -1, 450,
		-1, -1, -1, 750, -1, -1, -1, 960, 980, -1,
	}
	t.Run("Place in DrawState", func(t *testing.T) {
		game := PlaceItGame{
			Slots: slots,
			State: DrawState,
		}
		err := game.Place(9)

		checkType(t, err, "*game.ErrGameInvalidState")
	})
	t.Run("Place in PlaceState", func(t *testing.T) {
		game := PlaceItGame{
			Slots:      slots,
			State:      PlaceState,
			CurDraw:    145,
			ValidSlots: []int{3, 4},
		}
		err := game.Place(3)

		checkEquals(t, err, nil)
		checkEquals(t, game.State, DrawState)
		checkEquals(t, game.CurDraw, -1)
	})
	t.Run("Place Out of Bounds", func(t *testing.T) {
		game := PlaceItGame{
			Slots:   slots,
			State:   PlaceState,
			CurDraw: 145,
		}
		err := game.Place(100)

		checkNotNil(t, err)
		if err != nil {
			checkType(t, err, "*game.ErrPlaceOutofBounds")
		}
	})
	t.Run("Place in Invalid Slot", func(t *testing.T) {
		game := PlaceItGame{
			Slots:   slots,
			State:   PlaceState,
			CurDraw: 145,
		}
		err := game.Place(1)

		checkNotNil(t, err)
		if err != nil {
			checkType(t, err, "*game.ErrPlaceOutofBounds")
		}
	})
}

func TestValidSlots(t *testing.T) {
	slots := []int{
		-1, -1, 120, -1, -1, 200, -1, -1, -1, 450,
		-1, -1, -1, 750, -1, -1, -1, 960, 980, -1,
	}
	t.Run("Test Middle Draw", func(t *testing.T) {
		middleDraw := PlaceItGame{
			Slots:   slots,
			State:   PlaceState,
			CurDraw: 145,
		}
		got := validSlots(&middleDraw)
		want := []int{3, 4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Test Leading Edge Draw", func(t *testing.T) {
		leadEdgeDraw := PlaceItGame{
			Slots:   slots,
			State:   PlaceState,
			CurDraw: 15,
		}
		got := validSlots(&leadEdgeDraw)
		want := []int{0, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Test Tail Edge Draw", func(t *testing.T) {
		tailEdgeDraw := PlaceItGame{
			Slots:   slots,
			State:   PlaceState,
			CurDraw: 990,
		}
		got := validSlots(&tailEdgeDraw)
		want := []int{19}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Test Losing Draw", func(t *testing.T) {
		losingDraw := PlaceItGame{
			Slots:   slots,
			State:   PlaceState,
			CurDraw: 975,
		}
		got := validSlots(&losingDraw)
		var want []int

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Test Empty Draw", func(t *testing.T) {
		emptySlots := []int{
			-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
			-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
		}
		emptyDraw := PlaceItGame{
			Slots:   emptySlots,
			State:   PlaceState,
			CurDraw: 515,
		}
		got := validSlots(&emptyDraw)
		want := []int{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Test single number left", func(t *testing.T) {
		singleSlots := []int{
			-1, -1, -1, -1, -1, -1, -1, -1, -1, 500,
			-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
		}
		singleDraw := PlaceItGame{
			Slots:   singleSlots,
			State:   PlaceState,
			CurDraw: 150,
		}
		got := validSlots(&singleDraw)
		want := []int{
			0, 1, 2, 3, 4, 5, 6, 7, 8,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Test single number right", func(t *testing.T) {
		singleSlots := []int{
			-1, -1, -1, -1, -1, -1, -1, -1, -1, 500,
			-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
		}
		singleDraw := PlaceItGame{
			Slots:   singleSlots,
			State:   PlaceState,
			CurDraw: 750,
		}
		got := validSlots(&singleDraw)
		want := []int{
			10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestCheckGameover(t *testing.T) {
	t.Run("Test Game Win", func(t *testing.T) {
		winSlots := []int{
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
			31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		}
		winningGame := PlaceItGame{
			Slots:      winSlots,
			State:      PlaceState,
			CurDraw:    111,
			ValidSlots: []int{},
		}
		got := checkGameover(&winningGame)
		want := WinState

		checkEquals(t, got, want)
	})
	t.Run("Test Game Lose", func(t *testing.T) {
		loseSlots := []int{
			90, -1, -1, -1, -1, -1, -1, -1, -1, -1,
			-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
		}
		losingGame := PlaceItGame{
			Slots:      loseSlots,
			State:      PlaceState,
			CurDraw:    5,
			ValidSlots: []int{},
		}
		got := checkGameover(&losingGame)
		want := LoseState

		checkEquals(t, got, want)
	})
	t.Run("Test Not Gameover", func(t *testing.T) {
		normalSlots := []int{
			90, -1, -1, -1, -1, 500, -1, -1, -1, -1,
			-1, -1, -1, -1, -1, -1, -1, 901, -1, -1,
		}
		normalGame := PlaceItGame{
			Slots:      normalSlots,
			State:      PlaceState,
			CurDraw:    111,
			ValidSlots: []int{1, 2, 3, 4},
		}
		got := checkGameover(&normalGame)
		want := PlaceState

		checkEquals(t, got, want)
	})
}

func checkType[T any](t testing.TB, got T, want string) {
	t.Helper()
	gotType := reflect.TypeOf(got).String()
	if reflect.TypeOf(got).String() != want {
		t.Errorf("got variable of type %q want variable of type %q", gotType, want)
	}
}

func checkEquals[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func checkNotNil[T comparable](t testing.TB, got T) {
	t.Helper()
	var zero T
	if got == zero {
		t.Errorf("expected %v to not be nil", got)
	}
}
