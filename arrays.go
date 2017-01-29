package main

import "fmt"

func main() {

	var array [5]int
	array[0] = 4
	array[1] = 8
	fmt.Println("array", array)

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

}
