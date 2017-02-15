package main
import ("net/http"
	"strings"
	"fmt"
	"os"
	"io"
	"time"
	//"io/ioutil"
	"io/ioutil"
)

//fetch dump the html body of an url to stdout only if status code is 200
func fetch(ul string){
	if !strings.HasPrefix(ul, "http://"){
		ul = "http://" + ul
	}

	rep, err := http.Get(ul)
	if err != nil{
		fmt.Fprintf(os.Stderr, "fetch %v\n", err)
		os.Exit(1)
	}
	if rep.StatusCode != 200{
		fmt.Fprintf(os.Stderr, "%s\n", rep.Status)
		os.Exit(1)
	}

	// instead of using a large buffer(the commented part)
	// we can use io.Copy
	io.Copy(os.Stdout, rep.Body)
	//b, err := ioutil.ReadAll(rep.Body)
	//
	//rep.Body.Close()
	//
	//if err != nil{
	//	fmt.Fprintf(os.Stderr, "read %v\n", err)
	//	os.Exit(1)
	//
	//}
	//fmt.Printf("%v\n", string(b))

}

// in each goroutine we send processing results to the main goroutine
func fetch_one(ul string, ch chan <- string){
	if !strings.HasPrefix(ul, "http://"){
		ul = "http://" + ul
	}
	start := time.Now()
	rep, err := http.Get(ul)
	if err != nil{
		ch<- fmt.Sprintf("while accessing %v, %v",ul, err)
		return
	}
        // ignore the contents
	nbytes, err := io.Copy(ioutil.Discard, rep.Body)
	rep.Body.Close()
	if err != nil{
		ch <- fmt.Sprintf("while reading %v happended", err)
		return

	}
        elapsed := time.Since(start).Seconds()
	ch <- fmt.Sprintf("elapsed time: %v, url: %v, bytes: %v",elapsed,ul,nbytes)


}
func fetchall(){
	start := time.Now()
	urls := os.Args[1:]
	ch := make(chan string)
	for _, url := range urls{
		go fetch_one(url, ch)
	}

	for range os.Args[1:]{
		fmt.Println(<-ch)

	}
	fmt.Printf("elapsed: %.2fs", time.Since(start).Seconds())



}


func main() {
	//fetch("gopl.io")
	fetchall()
}
