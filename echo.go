package main
import ("fmt"
	"os"
        "strings"
	"time")

func echo1(){

       var s string
	for _,value := range os.Args[1:]{

	       s += value + " "

       }
	//fmt.Printf("%v\n",s)

}


func echo2(){

       var s string
	for i:=1;i<len(os.Args);i++{

	       s += os.Args[i] + " "

       }
	//fmt.Printf("%v\n",s)

}



func echo3(){


	strings.Join(os.Args[1:]," ")
	//fmt.Printf("%v\n",s)

}

func benchmark(f func()) (result string){

	start := time.Now()
	for i:=0;i<=1000000;i++{
		f()
	}

	elapsed := time.Since(start)
	result = fmt.Sprintf("%v:%v\n",f,elapsed)
	return
}

func main(){

        // it takes the last function (which uses string.join)
	// nearly half time as the first two to finish 1 million string concatenation ops
   	fmt.Println(benchmark(echo1))
	fmt.Println(benchmark(echo2))
	fmt.Println(benchmark(echo3))


}





