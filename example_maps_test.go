// This example is for a list of full names and
// how to extract the first name from it using Map.
package gopi_test

import (
	"fmt"
	"strings"

	. "github.com/mentix02/gopi"
)

func removeLastName(name interface{}) interface{} { return strings.Split(name.(string), " ")[0] }

func Example_map() {
	names := []interface{}{"Manan Yadav", "Aryan Verma", "Bruce Wayne"}
	for _, firstName := range Map(removeLastName, names) {
		fmt.Println(firstName)
		// do something with firstName
	}
}
