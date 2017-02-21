package main

import "fmt"

func main() {

	var array [5]int
	array[0] = 4
	array[1] = 8
	fmt.Println("array", array)

	// I can also use b:=[...]string{"H","e","l","l","o"}
	b := [5]string{"H", "e", "l", "l", "o"}

	for i := 0; i < len(b); i++ {

		//getting an element through its index
		fmt.Print(b[i])
	}

	fmt.Print("\n")

	var twodarray [2][2]float32

	for i := 0; i < 2; i++ {

		for j := 0; j < 2; j++ {

			twodarray[i][j] = float32(i + j)
			fmt.Println("i:", i, "j:", j)
		}

	}

	fmt.Println("2 dimentional dArray:", twodarray)

	double_array1 := [2][4]int {[4]int{1,2,3,4}, [4]int{5,6,7,8}}
	double_array2 := [2][4]int {[...]int{1,2,3,4}, [...]int{5,6,7,8}}
	double_array3 := [2][4]int {{1,2,3,4}, {5,6,7,8}}
	// for two arrays to be considered the same, they must have the same length
	// and elements type, [9]int != [10]int, once an array has been defined
	// you can not expand or shrink its size
	if (double_array1 == double_array2) && (double_array2 == double_array3){

		println("yes, they are all the same")
	}


	// specify an slice or array by index-value pair, for those not specified zero value
	// of its type, empty string here, indices can be out of order, can be omitted
	var color = []string {0:"red", 1:"black",  3:"pink", 5:"saphire", 2:"yellow", 20:"blue"}
	for i := range color{
		fmt.Printf("%d---->%v\n",i, color[i])
	}

}
