package main

import (
	"fmt"
)

// a slice parameter is just an alias of the original
// whatever you do with the parameter will visible to the
// caller
func reverse(a []byte){
	for i:=0; i<len(a)/2; i++{
		a[i],a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}
}

func reverse_test(){

	 b := []byte("0123456789")
	reverse(b)
	fmt.Println(string(b))
	b = []byte("123456789")
	reverse(b)
	fmt.Println(string(b))
}

func main() {

	a := make([]int, 3)
	a[0] = 2222

	//a = append(a,[2, 34])
	a = append(a, 2, 34)
	l := []int{100, 200}
	//note: use ... to expand l to 2 arguments
	a = append(a, l...)
	fmt.Println("after append a slice [100,200]", a)

	//slice assignment is not possible so far
	//a[1:3] = []int{3, 3}
	//fmt.Println("a:", a)
	aa := a[:3]
	fmt.Println("a[:3]", aa)
	aaa := make([]int, 2)

	//Note that you can not using negative indexes such as
	//aaa = a[:-2]
	aaa = a[:2]
	fmt.Println("a[:1]", aaa)

	//slice can be with variable-length, if you put a number in square bracket
	//it becames a 3-element array definition
	b := []string{"alex", "helen", "nova"}
	c := []string{}
	//c is empty and has zero elements
	//copy basically is the same as c = b[:len(c)]
	copy(c, b)
	fmt.Println("copy of b:", c)

	c = b[:len(c)]
	fmt.Println("slicing of b == equivalent of copy", c)

	//two-dimensional slice with variable-inner length
	d := make([][]int, 3)

	for i := 0; i < 3; i++ {

		innerlen := i + 1

		d[i] = make([]int, innerlen)

		for j := 0; j < innerlen; j++ {

			d[i][j] = i + j

		}

	}

	fmt.Println("two dimentinal slice with variable inner length", d)
	fmt.Println("len(d):", d)
	//can not do the following initiation
	//x := []byte{"a", "b", "c"}
	x := []byte("abc")
	fmt.Println("bytes array:", x)

	y := string(x)
	fmt.Println("convert bytest array to string:", y)

	u := "你好吗abc"
	v := []rune(u)
	fmt.Println("rune slice(convert string to character slice):", v)
	v[0] = 69
	w := string(v)
	fmt.Println("convert rune slice back to string:", w)

	slice1 := []string{"apple", "pear", "lemon"}
	slice2 := []string{"1", "1", "1", "1"}
	//this will copy the corresponding values to the target slice
	copy(slice2, slice1)
	fmt.Println("slice2:", slice2)

//=======================================================================

	// Declare an array of 10 bytes (ASCII characters). Remember: byte is uint8.
	var array [10]byte =[...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	//var array [...]byte =[...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	// Declare a and b as slice of bytes
	var a_slice, b_slice []byte

	// Make the slice "a_slice" refer to a "portion" of "array".
	a_slice = array[2:5]
	// a_slice is now a portion of array that contains: array[2], array[3] and array[4]

	// Make the slice "b_slice" refer to another "portion" of "array".
	b_slice = a_slice[1:]
	// b_slice is now a portion of array that contains: array[3], array[4]
	// pay attention to the capacity
        fmt.Printf("%p, cap: %v\n", &array[3], cap(array))
	fmt.Printf("%p, cap: %v\n", &a_slice[1],cap(a_slice))
	fmt.Printf("%p, cap: %v\n", &b_slice[0], cap(b_slice))

	a_slice = a_slice[1:]
	//notice that the difference between &a_slice[0] and &a_slice
	fmt.Printf("%p, cap: %v\n", &a_slice[0], cap(a_slice))
	fmt.Printf("%p, cap: %v\n", &a_slice, cap(a_slice))
//=======================================================================
        reverse_test()
}
