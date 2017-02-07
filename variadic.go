package main

import (
	"fmt"
)

// functions that accept a variable number of input parameters of the same type.
// They’re called variadic functions.
func custom_append(s []int, apd ...int) []int{

	for _, value :=  range apd {
		s = append(s, value)
	}
	return  s


}

func main() {

	var  original = []int{1,1,1,1,1,1,1}
	var  more = []int{2,2,2,2,2,2,2}
	// note that here if is a slice using ... after the slice parameter
	// just like append
	original = custom_append(original, more...)
	fmt.Println(original)
	original = custom_append(original, 3,3,3,3,3,3,3)
	fmt.Println(original)

}

// Note: for functions that take or return any type argument
// use interface {}