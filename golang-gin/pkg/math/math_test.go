package math

import "testing"

func TestAdd(t *testing.T) {
	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(6, 2)
	want := 4

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

