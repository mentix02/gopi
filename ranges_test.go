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

func helperRangeTest(start, end, step int, t *testing.T) []int {
	counter := start
	r := Range(start, end, step)
	for _, value := range r {
		if value != counter {
			t.Errorf("%d != %d", value, counter)
		}
		counter += 2
	}
	return r
}

func TestRangeEvenNums(t *testing.T) {
	helperRangeTest(-4, 10, 2, t)
}

func TestRangeWithWrongParams(t *testing.T) {
	r := helperRangeTest(5, 2, 1, t)
	if r != nil {
		t.Errorf("r is not nil.")
	}
}

func TestRangeOddNums(t *testing.T) {
	helperRangeTest(-5, 10, 2, t)
}
