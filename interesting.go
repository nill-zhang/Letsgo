package main

import "fmt"

type MyReader struct{}

func (r *MyReader) Read(b *[]byte) int {

	var c *[]byte
	/*c is useless here, just for fun*/
	c = b
	l := len(*c)
	fmt.Println("l:", l)
	for i := 0; i < l; i++ {

		/*append b itself everytime actually, b is changing everytime Read is called*/
		*c = append(*c, (*c)[i])
	}

	/*make it always reflect the newly append data*/
	*c = (*c)[l:]
	return l

}

func main() {
	r := &MyReader{}
	//f := &[]byte{0,0,0,0,0,0,0,0}
	f := make([]byte, 8)
	for o := 1; o < 10; o++ {

		/*let f be different, everytime we call c.Read*/
		for i := range f {

			f[i] = byte(o)

		}
		n := r.Read(&f)
		fmt.Printf("n:%v,f:%v\n", n, f)
	}
}
