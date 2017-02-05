package main

import ("fmt"; "unsafe")

type employee struct {
	id, name, address string

	salary float32
}

// Myint is just an alias for int32
type Myint int32



type Skills []string
type Human struct {
    	name string
    	age int
    	weight int
}

// When an anonymous field is a struct its fields are inserted
// (embedded) into the struct containing it
type Student struct {
    	//an anonymous field of type Human(no field name)
    	Human
	speciality string
	Skills //anonymous field for his skills
        int //we will use this int as an anonymous field for his preferred number
}

func (h *Human) SayHi() {
    fmt.Printf("Hi, I am %s, I am %v years-old\n", h.name, h.age)
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

	var  t Myint  = Myint(5)
	fmt.Println(t)
	fmt.Printf("type: %T\n",t)
	fmt.Println("size: ", unsafe.Sizeof(t))



	John := Student{Human:Human{"Mark", 25, 120},speciality:"Computer Science"}
	fmt.Printf("%v---%v---%v---%v\n",John.age, John.name, John.speciality,John.weight)
	// an anonymous field can be used with its type as its name!
	John.Human = Human{"Shaofeng", 31, 143}
	fmt.Printf("%v---%v---%v---%v\n",John.age, John.name, John.speciality,John.weight)
	fmt.Printf("%v---%v---%v---%v\n",John.Human.age, John.Human.name,
		                         John.speciality,John.Human.weight)
	John.Skills = []string{"C++", "Python", "Golang"}
	John.int =2222
	fmt.Printf("%v---%v\n",John.Skills, John.int)

	// If an anonymous field implements a given method, this method will be available
	// for the type that is using this anonymous field.
	John.SayHi()



}
