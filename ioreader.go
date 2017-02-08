package main
import ("fmt";"io";"strings";"time")


type Emit struct{

	s string
	num int
}

func (e Emit)Reader(p []byte) (n int, err error){

	if len(e.s) != 1{
		return 0, io.EOF
	}
	n = copy(p, strings.Repeat(e.s, e.num))
        return
}


func test_ioReader(){

	r := strings.NewReader("This is long input, lets do some thing一些中文支持吗 about it")
	b := make([]byte, 10)

	for {
		_,err := r.Read(b)
		if err == io.EOF { break }
		fmt.Printf("Data Read from Source :%v\n", string(b))
	}

}

func test_Custom_ioReader(){

	a := Emit{"A", 10}
	b:= make([]byte, 8)
	for{

		n,_ := a.Reader(b)
		fmt.Println("Data Read:", string(b[:n]))
		time.Sleep(time.Duration(time.Second*1))

	}



}
func main() {

	test_Custom_ioReader()
	test_ioReader()

}
