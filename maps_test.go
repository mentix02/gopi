package gopi

import (
	"strconv"
	"strings"
	"testing"
)

func square(n interface{}) interface{} { return n.(int) * n.(int) }

func removeLastName(name interface{}) interface{} { return strings.Split(name.(string), " ")[0] }

func convertToInt(num interface{}) interface{} {
	n, _ := strconv.Atoi(num.(string))
	return n
}

func TestSquareMap(t *testing.T) {
	x := []interface{}{2, 3, 4, 1}
	for index, value := range Map(square, x) {
		if square(x[index]) != value {
			t.Errorf("square(%v) != %v", square(x[index]), value)
		}
	}
}

func TestLastNameMap(t *testing.T) {
	firstNames := []string{"Manan", "Aryan", "Kunal"}
	names := []interface{}{"Manan Yadav", "Aryan Verma", "Kunal Pahuja"}
	for index, value := range Map(removeLastName, names) {
		if firstNames[index] != value {
			t.Errorf("removeLastName(%v) != %v", removeLastName(names[index]), value)
		}
	}
}

func TestSelfMap(t *testing.T) {
	nums := []interface{}{"2", "-41", "95"}
	intNums := []int{2, -41, 95}
	SelfMap(convertToInt, nums)
	for index, value := range nums {
		if intNums[index] != value {
			t.Errorf("%v != %v", intNums[index], value)
		}
	}
}
