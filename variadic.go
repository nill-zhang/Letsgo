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

func Max(lst ...int32) int32{
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

func Sum(lst  ...int32) int32{
	total := int32(0)
	for i:=0;i<len(lst);i++{
		total += lst[i]

	}
	return  total
}

// note that the following commented func declaration is wrong
// ... can only be used as the final parameter
// func Join(strs ...string, sep string)string{
func Join(sep string, strs ...string) (res string){
	if len(strs) == 0{
		res = sep
		return
	}
	res = ""
	for _, value := range strs[:len(strs) - 1] {
		res += value + sep
	}
	res =  res + strs[len(strs)-1]
	return
}


func Join_test(){
	p := fmt.Println
	p(Join("|","Alex","is","a","good","boy!"))
	p(Join("|",[]string{"Alex","is","a","good","boy!"}...))
	p(Join("",[]string{"Alex","is","a","good","boy!"}...))
	p(Join("",[]string{}...))
	p(Join("|",[]string{}...))
	p(Join("",[]string{"Oneword"}...))
}

func ArithFunc_test(){
	var list = []int32{2,2,3,12,-11,-2,232,221,102}
	// behaves like python sequence unpacking
	fmt.Printf("Sum(%v) = %v\n", list, Sum(list...))
	fmt.Printf("Min() = %v\n",Min(2,2,2,-1,0,-23,2323))
	fmt.Printf("Max() = %v\n",Max(2,2,2,-1,0,-2343,2323))
}

func CustomAppend_test(){
	var  original = []int{1,1,1,1,1,1,1}
	var  more = []int{2,2,2,2,2,2,2}
	// note that here if is a slice using ... after the slice parameter
	// just like append
	original = custom_append(original, more...)
	fmt.Println(original)
	original = custom_append(original, 3,3,3,3,3,3,3)
	fmt.Println(original)
}
func main() {

	Join_test()
	//ArithFunc_test()
	//CustomAppend_test()
}

// Note: for functions that take or return any type argument
// use interface {}