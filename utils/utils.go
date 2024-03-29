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

func Average(s []int) float32 {
	total := 0
	for _, v := range s {
		total += v
	}
	return float32(total) / float32(len(s))
}

func Histogram(s []int, max int) []int {
	retVal := make([]int, max)
	for i := range retVal {
		retVal[i] = 0
	}
	for _, v := range s {
		retVal[v-1]++
	}
	return retVal
}

func MaxSlicei(s []float64) int {
	if len(s) == 0 {
		return -1
	}

	maxIdx := 0
	max := s[maxIdx]
	for i, v := range s {
		if v > max {
			max = v
			maxIdx = i
		}
	}

	return maxIdx
}
