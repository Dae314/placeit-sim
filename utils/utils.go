package utils

import (
	"reflect"
	"testing"
)

func CheckType[T any](t testing.TB, got T, want string) {
	t.Helper()
	gotType := reflect.TypeOf(got).String()
	if reflect.TypeOf(got).String() != want {
		t.Errorf("got variable of type %q want variable of type %q", gotType, want)
	}
}

func CheckNotNil[T comparable](t testing.TB, got T) {
	t.Helper()
	var zero T
	if got == zero {
		t.Errorf("expected %v to not be nil", got)
	}
}

func CheckEquals[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func CheckContains[T comparable](t testing.TB, got []T, want T) {
	t.Helper()
	if !Contains(got, want) {
		t.Errorf("%v does not contain %v", got, want)
	}
}

func RemoveFromSlice(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func Contains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
