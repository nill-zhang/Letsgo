package main

import ("fmt"
        "io/ioutil"
        "bufio"
        "os"
	"strings"
)

func duplicate1(){

	linecount := make(map[string]uint16)
	line := bufio.NewScanner(os.Stdin)
	// scan next line, if true, extract its content
	for line.Scan(){
		linecount[line.Text()]++
	}

	Print_duplicate(linecount)




}

func duplicate2(){


	linecount := make(map[string]uint16)

	if len(os.Args) == 1{
		countline(os.Stdin, linecount)

	}else {
		for _, file := range os.Args[1:]{
			if f,err := os.Open(file); err == nil{
				countline(f, linecount)
				f.Close()
			}else{
				fmt.Printf("%v", err)
			}



		}
	}

	Print_duplicate(linecount)




}


func duplicate3(){

	linecount := make(map[string]uint16)
	if len(os.Args) == 1{
		fmt.Printf("no files to process\n")
	}else {
		for _, file := range os.Args[1:]{
			if b,err := ioutil.ReadFile(file); err != nil{
				continue
			}else{
				// b is of type byte[], with whole content
				for _, line := range strings.Split(string(b),"\n"){

					linecount[line]++
				}
			}

		}
	}

	Print_duplicate(linecount)
}

func Print_duplicate(dict map[string]uint16){

	for l,c := range dict{
		if c > 1{
			fmt.Printf("Count: %v\tLine: %v\n", c, l)
		}

	}
}

func countline(f *os.File, dict map[string]uint16){
       	line := bufio.NewScanner(f)
	for line.Scan(){
		dict[line.Text()]++
	}



}

func main() {

	//duplicate1()
	duplicate3()
}
