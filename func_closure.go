package main

import "fmt"

func adder() func(int) int {

	sum := 0

	return func(x int) int {

		sum += x
		return sum

	}

}

func main() {

	a, b := adder(), adder()

	for i := 0; i < 10; i++ {

		fmt.Println(a(i), "-----", b(i*i))

	}

}

/*
a,b func has its own sum which is inherited from where they were defined
their sum is independant of each other, and are attached to each function


results:

0 ----- 0
1 ----- 1
3 ----- 5
6 ----- 14
10 ----- 30
15 ----- 55
21 ----- 91
28 ----- 140
36 ----- 204
45 ----- 285

*/
