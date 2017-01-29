package main

import "fmt"

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
	//copy basiclly is copy c = b[:len(c)]
	copy(c, b)
	fmt.Println("copy of b:", c)

	c = b[:len(c)]
	fmt.Println("slicing of b == equivalent of copy", c)

	//two-dimentional slice with variable-inner
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
	//can not do the folloing initiation
	//x := []byte{"a", "b", "c"}
	x := []byte("abc")
	fmt.Println("bytes array:", x)

	y := string(x)
	fmt.Println("convert bytest array to string:", y)

	u := "good"
	v := []rune(u)
	fmt.Println("rune slice(convert stirng to character slice):", v)
	v[0] = 69
	w := string(v)
	fmt.Println("convert rune slice back to string:", w)

	slice1 := []string{"apple", "pear", "lemon"}
	slice2 := []string{"1", "1", "1", "1"}
	//this will copy the corresponding values to the target slice
	copy(slice2, slice1)
	fmt.Println("slice2:", slice2)

}
