package main

import (
	"fmt"
	"runtime"
	"time"
	"strings"
	"os"
)

var x, y = 3, 4

func DeferSimple(){

	var m, n = 10, 20

	//the arguments are assigned here and pushed to stack, code execution happens later
	//the second defer will be executed first because of Last in First out of Stack
	defer fmt.Println("After self-addition,x,y = ", x, y)
	defer fmt.Println("After self-substruction,m,n = ", m, n)

	//whatever follows defer will be executed first, when main is done, then execute defer statement
	fmt.Println("Before self-addition, x,y = ", x, y)
	fmt.Println("Before self-substruction,m,n = ", m, n)

	x++
	y++
	m--
	n--

	fmt.Println("x,y:", x, y)
	fmt.Println("m,n:", m, n)


}

func DeferMany(){
	fmt.Println("Counting Started")

	for i := 0; i < 5; i++ {


		// the function and argument expressions are evaluated when the statement
		// is executed, but the actual call is deferred until the function that contains
		// the defer statement has finished
		defer fmt.Println("i:", i)
		fmt.Printf("After Defer\n")

	}

	fmt.Println("Counting Finished")
}


func Trace(){
        start := time.Now()
	pc,_,_,ok := runtime.Caller(0)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		fmt.Printf("Entering function %v\n",details.Name())
	}
	// Alternative way:
	// runtime.FuncForPC(reflect.ValueOf(Trace).Pointer()).Name()
	defer func(s time.Time){
		pc,_,_,ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			fmt.Printf("Time Spent:\t%v\n",time.Since(s))
			fmt.Printf("Exiting function %v\n",details.Name())
		}
	}(start)
	time.Sleep(time.Duration(3)*time.Second)


}

func TraceRewrite(){

	// note that evaluation first, call at last
	defer DoTrace()()
	time.Sleep(time.Duration(3)*time.Second)

}

func DoTrace() func() {
	start := time.Now()
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		fmt.Printf("Entering function %v\n", details.Name())
	}
	return func() {
		pc, _, _, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			fmt.Printf("Time Spent:\t%v\n", time.Since(start))
			fmt.Printf("Exiting function %v\n", details.Name())

			}
	}

}

func ChangeReturn() (result uint32){

	defer func(s uint32){
		// For anonymous function, it inherits its closure' variables
		// defer will evaluate all the parameters at first
		// result is 0 initially, and passed to s, s is 0
		// when anonymous func is executed after ChangeReturn() return statement
		// has updated result
		// result is 3
		// anonymous function updates result, so result will be 6
		// it changes the returned value by assigning a new value to result
		fmt.Printf("Result is %v\n", s)
		fmt.Printf("Result is %v\n", result)
		// update the returning value, which must be a named result variable
		result = result*2
	}(result)
	time.Sleep(time.Duration(3)*time.Second)
	return uint32(3)

}

func OpenFiles(files []string){

	DoFile := func(file string){
		f, err := os.Open(file)
		if err != nil{

			fmt.Printf("%v while open %v\n",err, file)
		}
		defer f.Close()
	}
	for _,file := range files{
		//  you can use the following commented code snippet
		//  but you risk running out of file descriptor
		//  because for-loop opens too many files and does not
		//  close them until OpenFiles returns
		DoFile(file)

		//  f, err := os.Open(file)
		//  if err != nil{
		//
		//	  fmt.Printf("%v while open %v\n",err, file)
		//  }
		//  defer f.Close()

	}



}


func main() {
        p := fmt.Println

	p(strings.Repeat("*", 180))
	DeferSimple()


	p(strings.Repeat("*", 180))
	DeferMany()

	p(strings.Repeat("*", 180))
	Trace()

	p(strings.Repeat("*", 180))
	TraceRewrite()

	p(strings.Repeat("*", 180))
	p(ChangeReturn())

}

