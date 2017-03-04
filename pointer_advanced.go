package main

import "fmt"

type User struct {
	Name string
}


func Modify_test(){
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


func SliceArray(){
	names := [...]string{"Alex","Helen","Nova","Feven"}
	samenames := names[:]
	subnames := names[0:2]
	//fmt.Printf("%p\n", names)
	fmt.Printf("%p\n", &names)
	//fmt.Printf("%v\n", &names)
	fmt.Printf("%p\n", &names[0])
	//fmt.Printf("%p\n", names[0])
	fmt.Printf("%v\n", samenames)
	fmt.Printf("%p\n", samenames)
	// note that the following two are different
	// &samenames is a pointer with a value of the address of a slice variable
	// when you use %p to format samenames, you are getting the address of the slice
	// the address of the very first element
	fmt.Printf("%p\n", &samenames)
	fmt.Printf("%v\n", &samenames)
	fmt.Printf("%p\n", &samenames[0])
	fmt.Printf("%p\n", subnames)
}


func ReferenceNormal(re *uint, nm uint ){
        // go pass arguments by values, re and nm are local variables passed with
	// the same values as outer variables c and b accordingly
	// the only difference is re's value is an address(of a)
	// nm's value is a int, when you change re and nm, you just
	// change the local variable's values, it has no effect on
	// the outer c and b, because they are local copies
	// *re and *c are the same because they points to the same underlying
	// a, you reassign a by changing its values on the address of that memory
	// location(*re = 888)
	// &re and &c are completely different, because they are different variables' addresses
        // see http://stackoverflow.com/questions/4938612/how-do-i-print-the-pointer-value-of-a-go-object-what-does-the-pointer-value-mea
	// for details

	fmt.Printf("Inner: %v, %v\n",re, nm)
	fmt.Printf("Inner: %p, %p\n",&re, &nm)
	*re = 888
	re = nil
	nm = 0
}

func ReferenceNormal_test(){

	var a,b uint = 8,8
	c := &a
	fmt.Printf("Before: %v, %v\n",c, b)
	fmt.Printf("Before: %p, %p\n",&c, &b)
	ReferenceNormal(c,b)
	fmt.Printf("After: %v, %v\n",c, b)
	fmt.Printf("After: %p, %p\n",&c, &b)
	fmt.Printf("After: %v\n",*c)
	fmt.Printf("After: %v\n",a)

}

func main() {
	Modify_test()
	ReferenceNormal_test()
	SliceArray()
}