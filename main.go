package main
import (
	"fmt"
)

var ans [][]interface{}

func Pop(nested interface{}, flag bool) interface{} {
	fmt.Println("Incoming value is ", nested)
	switch v := nested.(type) {
	case [][]interface{}:
		var temp []interface{}
		if flag {
			// Recursively pop the first element from the first inner slice
			newValue := Pop(v[0], flag)

			temp = []interface{}{v[1:], newValue}
		} else {
			// Recursively pop the first element from the first inner slice
			newValue := Pop(v[len(v)-1], flag)

			temp = []interface{}{v[:len(v)-1], newValue}
		}

		fmt.Println(temp)
		// Create a new slice with the remaining elements and the result of recursion which contains all the elements except our removed
		return temp
	case []interface{}:
		// Pop the first element from the slice. We are sure this will work because the []interface is triggered
		fmt.Println("This is a normal Slice ", v)
		if flag {
			return v[1:]
		} else {
			return v[:len(v)-1]
		}
	default:
		// For other types, return as is
		return nested
	}
}

func main() {
	/*
		In Go, when you use a composite literal like {2}, it is inferred to be of a specific type based on the context in which it is used. In the case of [][]interface{}{{2}, {3}, {4}, {5, 6}}, the inner slices {2}, {3}, {4} are inferred to be of type []int, not []interface{}. This is because the elements within each inner slice are integers, not interfaces.
	*/

	/*
		Doubt why  is "temp" allowed but when the same thing is put in the nested structure "ns" it cause error
	*/

	ns := [][]interface{}{

		{"John", 30},
		{"Alice", 27},
		{"Bob", true},
		{3.14, "pi"},
		{
			[]interface{}{2}, []interface{}{3}, []interface{}{4}, []interface{}{5, 6},
		},
	}
	temp := [][]interface{}{
		{2}, {3}, {4}, {5, 6},
	}

	ans := Pop(ns, true)
	fmt.Println("ans is ", ans)
	fmt.Println(temp)

}
