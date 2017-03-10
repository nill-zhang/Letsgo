package main

import (
	"os"
	"strings"
	"fmt"
)

// LowerToUpper type satisfies io.Writer interface
// It will store and convert the bytes from p slices
// to its upper case string form
type LowerToUpper string

func (s *LowerToUpper) Write(p []byte) (n int, err error){
	*s  +=  LowerToUpper(strings.ToUpper(string(p)))
	return len(p),nil
}

func LowerToUpper_test(){

	var s LowerToUpper
	// _,err := c.Write([]byte("Alexandra "))
	_,err := s.Write([]byte{'A','l','e','x','a','n','d','r','a',' '})
	if err != nil{
		fmt.Fprintf(os.Stderr,"%v\n",err)
	}
	//fmt.Printf(string(c))
	// in order to satisfy io.Writer interface, I must provide an pointer concrete type
	fmt.Fprintf(&s,"%v\n","is a good guy")
	fmt.Printf(string(s))

}

// ByteCounter has a Write method its counts
// how many bytes are written from P slices
// I can also implement RuneCounter, WordCounter, LineCounter
type ByteCounter uint32

func (c *ByteCounter) Write(p []byte) (n int,err error){
	*c += ByteCounter(len(p))
	return len(p),nil
}

func ByteCounter_test(){
	var c ByteCounter
	c.Write([]byte{23,23,23,23,65})
	fmt.Printf("%v\n",c)

	fmt.Fprintf(&c, "io.writer,%v", "awesome!")
	fmt.Printf("%v\n",c)

}




func main(){

	LowerToUpper_test()
	ByteCounter_test()



}
