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

func Min(lst ...int32) int32{
	if len(lst) < 1{
		return -1
	}
	min := lst[0]
	for i:=1;i<len(lst);i++{
		if min > lst[i]{
			min = lst[i]

		}
	}
	return min
}

func Max(lst ...float32) float32{
	if len(lst) < 1{
		return -1.0
	}
	max := lst[0]
	for i:=1;i<len(lst);i++{
		if max < lst[i]{
			max = lst[i]

		}
	}
	return max

}

func Sum(lst  ...uint32) uint32{
	total := 0
	for i:=0;i<len(lst);i++{
		total += lst[i]

	}
return  total
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