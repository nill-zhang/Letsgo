package main

import (
	"fmt"
	"math"

)

// define a function type
type pickup_func func( float64) bool

func pickup_odd(element float64) bool{


	if int(element)%2 == 1{
		return  true
	}else{
		return  false
		}

}

func pickup_positive(element float64) bool{

		if element > 0 {
			return  true
		}else{

			return  false
		}

}

// use the function type pickup_func
func filter (source []float64, f pickup_func) (result []float64){

	for _, value := range source{

		if f(value){
			// since it's defined, so direct return here
			result = append(result, value)
		}
	}
	// because I use named return arguments, its already defined
	// so I use a naked return
	return
}

func filter_factory(f func(float64) bool) func([]float64)([]float64, []float64){

	return func(s []float64)(yes, no []float64) {

		for _, value := range s {
			if f(value) {
				yes = append(yes, value)
			} else {
				no = append(no, value)
			}
		}
		return
	}
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
	return func(x, y int) int {
		return x + y
	}

}


func test_generator_compute(){
	//func assigned to a variable
	//we don't need to define a func name here, after the assignment, it will get a name from the vairable
	sqrt := func(x, y float64) float64 {
		return math.Sqrt(x * x + y * y)
	}

	fmt.Println("sqrt:", compute(sqrt))
	fmt.Println("power:", compute(math.Pow))

	b := func_generator()
	fmt.Println("sum result:", b(22222, 11111))
}

func test_filter(){
	var original = []float64{2.2, 2,32, 232, 1.22, -2, -3, 4, 2,43,5, 7,9,-3.4}
	fmt.Println("Odd elements:", filter(original, pickup_odd))
	fmt.Println("Positvie elements:", filter(original, pickup_positive))
}


func test_filter_factory(){
	var original = []float64{8, -9,32, 232, 67, -2, -3, 4, 2,43,5, 7,9,-3.4}
	positives, negatives := filter_factory(pickup_positive)(original)
	fmt.Printf("Positives: %v\nNegatives: %v\n", positives, negatives)

	odds, evens := filter_factory(pickup_odd)(original)
	fmt.Printf("Odds: %v\nEvens: %v\n", odds, evens)

}


func main() {

	// defer execute before the caller(main) returns
	defer func(){
		print("Finished testing!!!!\n")
	}()

	// function literal, a function without a name
	// is a anonymous function
	func (){

		print("Start testing!!!!\n")
	}()


	test_generator_compute()
	test_filter()
	test_filter_factory()


}





