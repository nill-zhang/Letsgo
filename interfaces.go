package main

import (
	"unsafe"
	"fmt"
)

type I interface {
	Size() uintptr
}

type P interface {
	Size_pointer() uintptr
}

type T struct {

	s string
}

type Myfloat64 float64

//type implements interface by defined its methods
func (t T) Size() uintptr{

	return  unsafe.Sizeof(t.s)
}

func (f Myfloat64) Size() uintptr{

	return  unsafe.Sizeof(f)

}

func (t *T) Size_pointer() uintptr{

	return  unsafe.Sizeof((*t).s)
}

func (f *Myfloat64) Size_pointer() uintptr {

	return unsafe.Sizeof(*f)

}

func main() {

//==============================================================================
	t := T{"shaofeng zhang"}
	f := Myfloat64(3.14)


        var i I
	i = t
	fmt.Printf("%v 's size: %v\n", i, i.Size())
	i = &t
	fmt.Printf("%v 's size: %v\n", i, i.Size())
	i = f
	fmt.Printf("%v 's size: %v\n", i, i.Size())
	i = &f
	fmt.Printf("%v 's size: %v\n", i, i.Size())

	var p P
        //p = t
	//fmt.Printf("%v 's size: %v\n", p, p.Size_pointer())
	p = &t
	fmt.Printf("%v 's size: %v\n", p, p.Size_pointer())

	//p = f
	//fmt.Printf("%v 's size: %v\n", p, p.Size_pointer())
	p = &f
	fmt.Printf("%v 's size: %v\n", p, p.Size_pointer())
//==============================================================================
	// the difference between x and y is that
	// x is a nill interface, no concrete type and value of that concrete type
	// y is a non-nill interface, but its value is nill(*T pointer default), concrete type is *T
	var  x I
	var z *T
	var  y P = z

	// The above two lines is not equivalent to the following two line
	// the above z is an empty pointer type default value is nill
	// the bellow z is an empty string type, default values is ""
	// you will get different y because of that
	// var z T
	// var  y P  = &z

	fmt.Printf("%v----%p----%T\n", x, x, x)
	fmt.Printf("%v----%p----%T\n", y, y , y)
//================================================================================

	// r is called empty interface
	var  r interface{}

	// any type can be assigned to r
	// because there is no required methods or whatsoever for those types to implement
	r = 43
	fmt.Printf("%v----%T\n", r, r)

	r = 3.34343
	fmt.Printf("%v----%T\n", r, r)

	r = "awesome go"
	fmt.Printf("%v----%T\n", r, r)













}

// conclusion:
// [1].
// if a interface's concrete method implementation requires a pointer reveiver,
// the interface value must be assigned pointer of the concrete type to successfully
// call the method, it doesn't matter for the value receiver
//
// [2].
// an interface value can be assigned by any concrete type, only if that type implements
// all the interface's methods
//
// [3].
// Don't call an method on an nill interface, because it does not know which concrete
// method to call since it does not have a underlying concrete type
