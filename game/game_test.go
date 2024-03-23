package game

import (
	"reflect"
	"testing"
)

func TestNewGame(t *testing.T) {
	game := NewGame()
	gameType := reflect.TypeOf(game).String()

	if gameType != "game.PlaceItGame" {
		t.Errorf("Expected game to be type game.PlaceItGame, but it is %q", gameType)
	}
}
