package main

import ("net/http"
           "fmt"
           "sync"
	    "log"
	"strings"
)

var count uint16
var mu sync.Mutex
func Index(w http.ResponseWriter, r *http.Request){

	Counter()
	fmt.Println("index executed")
	fmt.Fprintf(w, "your are accessing %v from %v\n", r.URL.Path, r.RemoteAddr)

}
func Counter(){

	// prevent two different user from updating the counter
	// at the same time
	mu.Lock()
	count++
	mu.Unlock()
}

func Count(w http.ResponseWriter, r *http.Request){
	Counter()
	fmt.Println("count executed")
	fmt.Fprintf(w, "%v visits to this website!!\n", count)


}

func Greet(w http.ResponseWriter, r *http.Request){
	Counter()
	fmt.Println("greet executed")
	// note that you must use Request.ParseForm first
	// before you can extract Form data
	if r.ParseForm() == nil {
		fn := r.Form.Get("firstname")
		ln := r.Form.Get("lastname")
		fmt.Fprintf(w, "Hi, %v %v!\n\n" +
				"The following are Request header information\n", fn,ln)
	}
	for k, v := range r.Header{
		fmt.Fprintf(w, "Header[%q]:\t\t%s\n",k,strings.Join(v, " "))
	}




}

func main() {

	http.HandleFunc("/", Index)
	http.HandleFunc("/count", Count)
	http.HandleFunc("/greet", Greet)
	log.Fatal(http.ListenAndServe("localhost:800", nil))

}
