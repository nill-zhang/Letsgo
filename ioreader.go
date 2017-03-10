package main
import ("fmt";"io";"strings";"time"

)


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


type MyStringReader struct{
	        str string
		cur int
	}

func(R *MyStringReader) Read(p []byte)(m int, err error){
		if len(p) + R.cur <= len(R.str) {
			m = copy(p, []byte(R.str[R.cur:len(p) + R.cur]))
			R.cur = R.cur + m
			err = nil
			return
		}
		m = copy(p, []byte(R.str[R.cur:]))
		err = io.EOF
		return
	}


func NewMyStringReader(s string) io.Reader{

	r := MyStringReader{s,0}
	return &r
}


func MyStringReader_test(){
	y := NewMyStringReader("This is freaking awesome!!")
	p := make([]byte, 4)
	for{
		n,err := y.Read(p)
		// you should not print n-p pair after if-statement
		// because you get out of the loop early and leave out the final
		// pair
		if n < len(p){
			fmt.Printf("%d bytes read, Buffer: %v\n",n,string(p[:n]))
		}else{
			fmt.Printf("%d bytes read, Buffer: %v\n",n,string(p))
		}
		if err != nil{
			break
	}

	}
}

type LimitReader struct{
	r io.Reader
	n int
        total int
	}

func (l *LimitReader)Read(p []byte)(m int, err error){


	m,err = l.r.Read(p)
	l.total += m
		if l.total > l.n {
			m = m-(l.total-l.n)
			// the following line of code does not work
			// even though you changed p here
			// the caller still sees the full p
			// p = p[:m]
			err = io.EOF
			return
		}
		return
}


func NewLimitReader(r io.Reader, n int)  *LimitReader{

	return &LimitReader{r,n,0}
}

func LimitReader_test(){

	p := make([]byte, 4)
	r := strings.NewReader("this is a test function which tests LimitReader")
	y := NewLimitReader(r,11)
	for{
		n,err := y.Read(p)
		// you should not print n-p pair after if-statement
		// because you get out of the loop early and leave out the final
		// pair
		if n < len(p){
			fmt.Printf("%d bytes read, Buffer: %v\n",n,string(p[:n]))
		}else{
			fmt.Printf("%d bytes read, Buffer: %v\n",n,string(p))
		}
		if err != nil{
			break
	}

	}


}

func main() {

	//test_Custom_ioReader()
	//test_ioReader()
	MyStringReader_test()
	LimitReader_test()

}
