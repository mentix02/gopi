package gopi_test

import (
	"fmt"
	. "github.com/mentix02/gopi"
)

func Example() {
	// An example of a positive range generator
	for _, value := range PosRange(10) {
		// do something with value
		fmt.Println(value)
	}
	// An example of all even numbers from [-4, 6)
	for _, num := range Range(-4, 6, 2) {
		// do something with num
		fmt.Println(num)
	}
}
