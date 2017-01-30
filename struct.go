package main

import "fmt"

type employee struct {
	id, name, address string

	salary float32
}

func main() {

	ep1 := employee{id: "1001", name: "Alex"}
	var p = &ep1

	fmt.Println("pointer address:", p)
	fmt.Println("id: ", p.id)
	ep1.address = "404 55 Oakmount Rd."
	fmt.Println("Employee1:", ep1)

	ep2 := employee{"1003", "Helen", "20 Humberline Drive", 4000.00}
	p = &ep2
	fmt.Println("pointer address:", p)
	p.name = "Helen Zhang"
	fmt.Println("employee2:", ep2)
	fmt.Println("*p(employee2):", *p)

	if ep2.salary >= 3000 {

		fmt.Println(ep2.name, "is rich")
	}
}
