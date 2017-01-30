package main

import (
	"fmt"
)

var x, y = 3, 4

func main() {

	var m, n = 10, 20

	//the arguments are assigned here and pushed to stack, code execution happens later
	//the second defer will be executed first because of Last in First out of Stack
	defer fmt.Println("After self-addition,x,y = ", x, y)
	defer fmt.Println("After self-substruction,m,n = ", m, n)

	//whatever follows defer will be executed first, when main is done, then execute defer statement
	fmt.Println("Before self-addition, x,y = ", x, y)
	fmt.Println("Before self-substruction,m,n = ", m, n)

	x++
	y++
	m--
	n--

	fmt.Println("x,y:", x, y)
	fmt.Println("m,n:", m, n)

	fmt.Println("Counting Started")

	for i := 0; i < 5; i++ {

		fmt.Print("Before Defer")
		//I add more defers, make it total seven, the last one will execute first
		defer fmt.Println("i:", i)
		fmt.Println("After Defer")

	}

	fmt.Println("Counting Finished")

}

//Results:

//////////////////////////////////////////////////////*
/*

Before self-addition, x,y =  3 4
Before self-substruction,m,n =  10 20
x,y: 4 5
m,n: 9 19
Counting Started
Before DeferAfter Defer
Before DeferAfter Defer
Before DeferAfter Defer
Before DeferAfter Defer
Before DeferAfter Defer
Counting Finished
i: 4
i: 3
i: 2
i: 1
i: 0
After self-substruction,m,n =  10 20
After self-addition,x,y =  3 4

*/
////////////////////////////////////////////////////////
