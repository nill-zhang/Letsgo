package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	m, n float64
}

func main() {

	var v Vertex= Vertex{3.5, 5.66}
	pt := &v
	fmt.Println(v.Abs_pointer_receiver())
	fmt.Println(pt.Abs_value_receiver())
	fmt.Println(v.Abs_value_receiver())
	fmt.Println(pt.Abs_value_receiver())
	fmt.Println((*pt).Abs_value_receiver())
	fmt.Println((*pt).Abs_pointer_receiver())


	//fmt.Println(Abs_pointer_argu(v)) invalid argument type
	fmt.Println(Abs_value_argu(v))
	fmt.Println(Abs_pointer_argu(pt))
	//fmt.Println(Abs_value_argu(pt)) invalid argument type

	var  t int32  = 10
	var  w Vertex  =Vertex{6,8}
	fmt.Printf("Before calling Scale, v: %v-----t: %v\n", w, t)
	(&w).Scale_shared_argu(&t)
	fmt.Printf("After calling Scale, v: %v-----t: %v\n", w, t)

	// I did not define new variables, just use the old ones
	fmt.Printf("Before calling Scale, v: %v-----t: %v\n", w, t)
	w.Scale_shared_argu(&t)
	fmt.Printf("After calling Scale, v: %v-----t: %v\n", w, t)


	fmt.Printf("Before calling Scale, v: %v-----t: %v\n", w, t)
	w.Scale_copied_argu(t)
	fmt.Printf("After calling Scale, v: %v-----t: %v\n", w, t)

	fmt.Printf("Before calling Scale, v: %v-----t: %v\n", w, t)
	(&w).Scale_copied_argu(t)
	fmt.Printf("After calling Scale, v: %v-----t: %v\n", w, t)


}

func (v *Vertex) Abs_pointer_receiver() float64 {

	return math.Sqrt(math.Pow(v.m, 2) + math.Pow(v.n, 2))

}

func (v Vertex) Abs_value_receiver() float64 {

	return math.Sqrt(math.Pow(v.m, 2) + math.Pow(v.n, 2))

}

func Abs_value_argu(v Vertex) float64{

	return  math.Sqrt(math.Pow(v.m, 2)+math.Pow(v.n, 2))
}

func Abs_pointer_argu(v *Vertex) float64{

	return  math.Sqrt(math.Pow(v.m, 2)+ math.Pow(v.n, 2))

}

func (v *Vertex) Scale_shared_argu(t *int32){

	v.m *= float64(*t) 
	v.n *= float64(*t)
	(*t) *= *t

}


func (v Vertex) Scale_copied_argu(t int32){

	v.m *= float64(t)
	v.n *= float64(t)
	t *= t
	

}


// Conclusion:
//[1].
// With Pointer receivers and pointer arguments, you can modify their data, the
// data is shared with the original data, but with value receiver, you get copied data
// whatever you do with the receiver and arguments within the method will not affect
// the original, because methods get a copied version, use pointers can save memory
// usually methods use pointers because they need to change the receiver's state or inner
// data and even though like Abs function, it doesn't change the receiver's original data
// it is better to still use the pointer receiver because we don't need to copy the
// whole struct data if it is very big
//
//[2].
// Did you notice that it doesn't matter whatever receiver you pass to the method(Scale_xxx)
// it will convert to it requires if necessary
