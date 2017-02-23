package main

import (
	"fmt"

	"os"
	"bufio"

	"unicode"
)

func map_test(){
	m := make(map[string]int)
	m["age"] = 23
	m["id"] = 1009
	m["height"] = 176

	fmt.Println(m)

	for i, j := range m {

		/*we use printf here because we are formating output
		and don't forget the \n at the end*/
		fmt.Printf("%s---->%d\n", i, j)

	}

	/*another way to define a map*/
	var m2 map[string]string
	m2 = make(map[string]string)
	m2["name"] = "Shaofeng zhang"
	m2["address"] = "404 55 Oakmount Road."


	fmt.Println(m2)

	for i, j := range m2 {

		fmt.Printf("%s---->%q\n", i, j)

	}

	slice := []string{"A", "B", "C", "D", "E"}

	/*ignore the index, fetch the items*/
	for _, item := range slice {

		fmt.Println(item)

	}

	/*only the index, this is useful for map type*/
	for idx := range slice {

		fmt.Println(idx)

	}

}

func WordFreq(file string){

	wordcount := make(map[string]int)
	f, success := os.Open(file)
	if success != nil{
		fmt.Fprintf(os.Stderr, "%v while open %v", success, file)
		os.Exit(1)
	}

	// note that we can not use
	// scanner :=  &bufio.Scanner{
	//                      r:f,
	//                      split:bufio.ScanWords,
	//                      MaxTokenSize: bufio.MaxScanTokenSize,
	//     }
	// to initialize our scanner because bufio.r bufio.split...
	// can not be exported, we need to  use its exported method
	// bufio.NewScanner method, if we want to change its split fun
	// we use the exported bufio.ScanWords
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for {
		if scanner.Scan(){
			wordcount[scanner.Text()]++
		}else{
			break
		}

        for i,j :=range wordcount{

		fmt.Printf("%v----%d\n",i,j)
	}

	}





}


func CharCount(file string){
        count := make(map[string]int)
	f, status := os.Open(file)
	if status != nil{
		fmt.Fprintf(os.Stderr, "can not open %v, error:%v", file,status)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)
	for {
		if scanner.Scan(){
			r := []rune(scanner.Text())[0]
			switch {
				case unicode.IsSymbol(r): count["symbol"]++
				case unicode.IsSpace(r): count["space"]++
				case unicode.IsNumber(r): count["number"]++
				case unicode.IsDigit(r): count["digit"]++
				case unicode.IsPunct(r): count["punct"]++
				case unicode.IsPrint(r): count["print"]++
				case unicode.IsControl(r): count["control"]++
			}

		}else{
			break
		}


	}

	for i,j :=range count{

		fmt.Printf("%v----%d\n",i,j)
	}

}



func main() {
	WordFreq("I:/learning/EnglishLearning/GRE/English+英文原着500本/English 英文原著shu/a message from the sea(大海来信).txt")
	CharCount("I:/learning/EnglishLearning/GRE/English+英文原着500本/English 英文原著shu/a message from the sea(大海来信).txt")

}
