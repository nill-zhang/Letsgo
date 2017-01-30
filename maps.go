package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int)
	m["age"] = 23
	m["id"] = 1009
	m["height"] = 176

	fmt.Println(m)

	for i, j := range m {

		/*we use printf here because we are formating output
		and don't forget the \n at the end*/
		fmt.Printf("%s---->%d\n", i, j)

	}

	/*another way to define a map*/
	var m2 map[string]string
	m2 = make(map[string]string)
	m2["name"] = "Shaofeng zhang"
	m2["address"] = "404 55 Oakmount Road."

	fmt.Println(m2)

	for i, j := range m2 {

		fmt.Printf("%s---->%q\n", i, j)

	}

	slice := []string{"A", "B", "C", "D", "E"}

	/*ignore the index, fetch the items*/
	for _, item := range slice {

		fmt.Println(item)

	}

	/*only the index, this is useful for map type*/
	for idx := range slice {

		fmt.Println(idx)

	}

}
