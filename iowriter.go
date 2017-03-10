package main

import (
	"os"
	"strings"
	"fmt"
	"io"
	"log"
	"io/ioutil"
	"path/filepath"
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

type CountingWriter struct{
	number *int64
	w io.Writer
}

func (c *CountingWriter) Write(p []byte)(n int, err error){
	n,err = c.w.Write(p)
	*(c.number) += int64(n)
	return

}

func NewCountingWriter(w io.Writer)(io.Writer, *int64){
	newW := &CountingWriter{new(int64),w}

	return newW,newW.number


}

func CountingWriter_test(){
        // change to the letsgo directory and create 2 absolute filepath
	os.Chdir("C:/Users/Admin/Documents/golang/resources/letsgo")
	rpath,err1 := filepath.Abs("defer.go")
	wpath,err2 := filepath.Abs("test.txt")
	if err1!= nil || err2 != nil{
		log.Fatalf("Can not create absolute path, (%v,%v)\n", err1, err2)
	}

	bytes,err := ioutil.ReadFile(rpath)
	if err != nil{
		log.Fatalf("%v while open %v\n", err, rpath)
	}
	fw, err := os.Create(wpath)
	defer fw.Close()
	if err != nil{
		log.Fatalf("%v while creating %v\n", err, wpath)
	}
	nw, number := NewCountingWriter(fw)
	n, _ := nw.Write(bytes)

	fmtstr := "%v bytes read, total %v bytes written\n"
	p:= fmt.Printf

	p(fmtstr,n,*number)

	n, _ = nw.Write([]byte("done!\n"))
	p(fmtstr,n,*number)

	n, _ = nw.Write([]byte("see you"))
	p(fmtstr,n,*number)

	n, _ = nw.Write([]byte("awesome job!!"))
	p(fmtstr,n,*number)
}




func main(){

	LowerToUpper_test()
	ByteCounter_test()
	CountingWriter_test()





}
