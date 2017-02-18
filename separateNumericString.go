package main

import (
	"fmt"
	"bytes"
	"strings"
)

func separateNumeric(in string) (out string){
	var head, tail string
	var sepNumber int
	var buf bytes.Buffer
	// check if input string is a float string
	if strings.Contains(in, "."){

		head = strings.Split(in, ".")[0]
		tail = strings.Split(in, ".")[1]
		in = head
	}

	remainder := len(in)%3
	step := len(in)/3
	// if input string's length is not divisible by 3, we need to write 1 or 2 remainder digits first
	// after that we then check, do we need to add a `,` separator because if we don't
	// strings like 1234.00 will be written as 1
	if remainder != 0 {
		buf.WriteString(in[:remainder])
		if step > 0 {
			buf.WriteString(",")
		}
	}

	// now lets dealing with the rest
	// after output the remainder, now the rest will be multiple 3 digits
	// for step == 0(``) or step == 1(`123`), we don't need to write any separator
	// for strings length no longer than 5, we need to start with index begin(begin can be 0,1,2)
	// for other strings with length longer than 5 we just write the first 3 digits
	// sepNumber stands for how many `,` separator we need to add thereafter
	// for strings with remainder == 0, we need step-1 separators, for example 123456789 we need two
	// for strings with remainder != 0, for strings like `1.23` `23452` `1234.22`, we need step separators

	sepNumber = step -1

	if sepNumber < 1 {
		buf.WriteString(in[remainder:])
	}else{
		buf.WriteString(in[:3])
		for i:= 1; i<=sepNumber;i++{
			buf.WriteString(",")
			buf.WriteString(in[i*3:i*3+3])


		}
	}
	// append the decimal point and the rest of the float string
	if tail != ""{
                buf.WriteString(".")
		buf.WriteString(tail)
	}

	// convert []byte to string
	out = buf.String()
	return

}

func separateNumeric_test(teststr string){

	fmt.Printf("Input: %v\tOutput: %v\n",teststr,separateNumeric(teststr))
}

func main() {
	separateNumeric_test("134567.3252323")
	separateNumeric_test("134567232")
	separateNumeric_test("134232467.3252323")
	separateNumeric_test("12134232467.3252323")
	separateNumeric_test("1134232467.3252323")
	separateNumeric_test("111134232467.3252323")
	separateNumeric_test("1111134232467.3252323")
	separateNumeric_test("1342.00")
	separateNumeric_test("13")
	separateNumeric_test("1.11")
}
