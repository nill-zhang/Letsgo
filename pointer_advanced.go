package main

import "fmt"

type User struct {
	Name string
}

func main() {
	u1 := &User{Name: "Leto"}
	u2 := User{Name: "Leto"}
	u3 := &User{Name: "Leto"}
	println("before modification: ", u1.Name)
	Modify1(u1)
	println("after modification: ", u1.Name)

	println("before modification: ", u2.Name)
	Modify2(u2)
	println("after modification: ", u2.Name)

	println("before modification: ", u3.Name)
	Modify3(&u3)
	println("after modification: ", u3.Name)

	sum := 0
	var double_sum *int //a pointer to int
	for i := 0; i < 10; i++ {
		sum += i
	}
	double_sum = new(int) //allocate memory for an int and make double_sum point to it
	*double_sum = sum * 2 //use the allocated memory, by dereferencing double_sum
	fmt.Println("The sum of numbers from 0 to 10 is: ", sum)
	fmt.Println("The double of this sum is: ", *double_sum)
}

func Modify1(u *User) {
	// u is an variable, pointer type
	// we change the pointer's value, didn't change the the value
	// of the type the pointer pointing to(*u = User{Name:"Paul"})
	// u is u1's local copy, they both have the some value(address of
	// the struct User{Name: "Leto"}), u1 is not affected by changing u
	// but will be affected if its underlying contents of that address
	u = &User{Name: "Paul"}

}

func Modify2(u User) {

	u.Name = "Duncan"
}

func Modify3(u **User) {

	// this one is tricky, u is a variable, a pointer
	// it has the value of the address of a pointer(u3)(&User{Name: "Leto"})
	// *u is that pointer(u3)'value (&User{Name: "Leto"})
	// we change(u3) its content to reference to another address(User{Name: "Bob"})

	*u = &User{Name: "Bob"}
}
