package main

import (
	"fmt"
	"unicode/utf8"
)

// a slice parameter is just an alias of the original
// whatever you do with the parameter will visible to the
// caller
func Reverse(a []byte){
	for i:=0; i<len(a)/2; i++{
		a[i],a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}
}

// looks complicated, it is very easy, rotate behaves like block shift
// cut it then move the two separated parts
func Rotate(s []int, step int)[]int {

	if step < 0 {
		s = append(s[-step:], s[:-step]...)
	} else {
		s = append(s[len(s) - step:], s[:len(s) - step]...)
	}
	return s
}

func RemoveAdjacentDuplicates(s []int) []int{
	// not using s[0], because we need a slice not an element
	b := s[:1]
	for _, value := range s[1:]{
		if value != b[len(b)-1]{
			b = append(b,value)}
	}

	return b
}

func ReverseCharacters(b []byte) []byte{
	/* I use a in-place technique to do the reverse
	   operation, I didn't allocate new momery, insted
	   , I shift the unprocessed slice to the end after I
	   extract the last rune and put it at the beginning
	*/
	a := b[:0]
	for {
		rune, size := utf8.DecodeLastRune(b)
		if rune == utf8.RuneError && size == 0{
			break
		}
		// just can not use the following
		// b[size:] := b[:len(b) - size]
		copy(b[size:],b[:len(b) - size])
		a = append(a, []byte(string(rune))...)
		b = b[size:]
	}
	return a


}


func Reverse_test(){

	b := []byte("0123456789")
	Reverse(b)
	fmt.Println(string(b))
	b = []byte("123456789")
	Reverse(b)
	fmt.Println(string(b))
}

func Rotate_test(){
	a := []int{3,4,5,6,7,8,9,10,11,12,13}
	fmt.Printf("a: %v\nleft rotate a by 2 spots: %v\n",a, Rotate(a, -2))
	fmt.Printf("a: %v\nright rotate a by 3 spots: %v\n",a, Rotate(a, 3))

}


func RemoveAdjacentDuplicates_test(){
	a := []int{-1, 2, -3, -2, -2, -2, -2, 3,2,3,90,87,4,4,0,5,2}
	fmt.Printf("a: %v\n",a)
	fmt.Printf("after removing duplicates: a: %v\n",RemoveAdjacentDuplicates(a))



}

func ReverseCharacters_test(){
	 comments := "张少锋的yingyu很好！"
	 newcomments := ReverseCharacters([]byte(comments))
	 fmt.Printf("comments is %v\nafter reverse its characters" +
		" comment is %v\n", comments, string(newcomments))



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
        Reverse_test()
	RemoveAdjacentDuplicates_test()
	ReverseCharacters_test()
	Rotate_test()
}
