package utils

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	testSlice := []int{1, 2, 3, 4}
	t.Run("Test for item in slice", func(t *testing.T) {
		got := Contains(testSlice, 4)
		want := true
		CheckEquals(t, got, want)
	})
	t.Run("Test for item not in slice", func(t *testing.T) {
		got := Contains(testSlice, 100)
		want := false
		CheckEquals(t, got, want)
	})
}

func TestRemoveFromSlice(t *testing.T) {
	testSlice := []int{1, 2, 3, 4}
	got := RemoveFromSlice(testSlice, 1)
	want := []int{1, 4, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestAverage(t *testing.T) {
	testSlice := []int{1, 2, 3, 4}
	got := Average(testSlice)
	want := float32(2.5)
	CheckEquals(t, got, want)
}
