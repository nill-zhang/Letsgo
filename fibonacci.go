package main

import "fmt"

func fibonacci(x int) int {

	if x >= 2 {

		return fibonacci(x-1) + fibonacci(x-2)

	} else if x >= 0 {

		return x
	}

	return 0
}

func main() {

	fmt.Println(fibonacci(9))

}
