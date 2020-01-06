// This example shows how to compute all the even
// numbers between -5 and 10 using Range.
package gopi_test

import (
	"fmt"
	. "github.com/mentix02/gopi"
)

func Example_range() {
	for _, value := range Range(-5, 10, 2) {
		fmt.Println(value)
		// do something with value
	}
}

func normalForLoopWay() {
	for value := -5; value < 10; value += 2 {
		fmt.Println(value)
		// do something with value
	}
}
