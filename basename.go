package main

import ("fmt"
	"strings"
)


// Note:
// finish this function seems easy
// but it takes me 40 minutes to finish it, the original
// code snippet is not working in the book. That's why we
// need to write our own while study them.
// the problem is that for each character in string
// its type is byte, we can not simply assert
// ul[i] == "/", because we are comparing two
// different types. we also can not write the comparision as
// ul[i] == byte("/"), it seems to work. but unfortunately
// the argument is a string, it may contain more than one
// character, although it has only one character in our situation.
// solution: we need to convert the "/" to slice of type byte(byte arrays)
// then do the comparison with its first element which is of type byte
func getBasename(ul string) (filename string){
	filename = ul
	// get the first path sep from right hand side
	for i:=len(ul)-1; i>=0; i--{
		if ul[i] == []byte("/")[0] {
			filename = ul[i+1:]
			break
		}

	}
        // for the filename get the first . from right hand side
	for j := len(filename)-1;j >= 0;j--{
              if filename[j] == []byte(".")[0]{
		      filename = filename[:j]
		      break
	      }

	}
	return
}

func getBasename_test(teststr string){
	fmt.Printf("Input: %q\tOutput: %q\n",teststr,getBasename(teststr))

}

func getBasename_short_test(teststr string){
	fmt.Printf("Input: %q\tOutput: %q\n",teststr,getBasename_short(teststr))

}

// use some packages to ease our life
func getBasename_short(ul string) (bname string){

	a := strings.LastIndex(ul, "/")
	b := strings.LastIndex(ul, ".")
	// adjust a = -1 to make a+1 start from beginning
	if a == -1{a = -1}
	if b == -1{b = len(ul)}
	bname = ul[a+1:b]
	return

}
func main() {
	getBasename_test("abc")
	getBasename_test("/etc/http/conf/http.conf")
	getBasename_test("/var/http/index.html.backup")
	getBasename_short_test("abc")
	getBasename_short_test("/etc/http/conf/http.conf")
	getBasename_short_test("/var/http/index.html.backup")

}
