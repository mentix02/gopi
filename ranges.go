package gopi

// PosRange only takes in one positive argument and
// returns a slice of all the natural numbers beginning
// from 0 upto n - 1 in an unsigned integer slice.
func PosRange(n int) []uint {
	var counter uint
	s := make([]uint, n)
	for i := range s {
		s[i] = counter
		counter++
	}
	return s
}

// Range is the most similar to Python's range as it
// allows for negative numbers as well range from
// [start, end - 1) with a step after each iteration.
func Range(start, end, step int) []int {
	if end <= start {
		return []int{}
	}

	diff := end - start

	var s []int

	if diff%2 == 0 {
		s = make([]int, diff/2)
	} else {
		s = make([]int, diff/2+1)
	}

	for value, index := start, 0; value < end; value, index = value+step, index+1 {
		s[index] = value
	}

	return s
}
