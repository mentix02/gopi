package gopi

import "testing"

func TestPosRange(t *testing.T) {
	var counter uint
	for _, value := range PosRange(10) {
		if value != counter {
			t.Errorf("%d != %d", value, counter)
		}
		counter++
	}
}

func TestRange(t *testing.T) {
	start, end, step := -5, 10, 2
	counter := start
	for _, value := range Range(start, end, step) {
		if value != counter {
			t.Errorf("%d != %d", value, counter)
		}
		counter += 2
	}
}
