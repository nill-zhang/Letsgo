package main

import (
	"fmt"
	"math"
)

func main() {

	//func assigned to a variable
	//we don't need to define a func name here, after the assignment, it will get a name from the vairable
	sqrt := func(x, y float64) float64 { return math.Sqrt(x*x + y*y) }

	fmt.Println("sqrt:", compute(sqrt))
	fmt.Println("power:", compute(math.Pow))

	b := func_generator()
	fmt.Println("sum result:", b(22222, 11111))

}

//take func as a argument
func compute(fn func(float64, float64) float64) float64 {

	return fn(4, 3)
}

//return a func
func func_generator() func(int, int) int {

	// add := func(x, y int) int { return x + y }
	// return add
	// the above two lines equals the following
	return func(x, y int) int { return x + y }

}
