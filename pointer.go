package main

import (
	"fmt"
)

func main() {

	var x int = 100
	var y string = "China"

	px := &x
	var py *string = &y

	fmt.Println("px: ", px)
	fmt.Println("px's contents: ", *px)

	fmt.Println("py: ", py)
	fmt.Println("py's contents: ", *py)

	*px = *px + 100
	*py = *py + "Anhui"

	fmt.Println("After modication")
	fmt.Println("px: ", px)
	fmt.Println("px's contents: ", *px)

	fmt.Println("py: ", py)
	fmt.Println("py's contents: ", *py)

}
