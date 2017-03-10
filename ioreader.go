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


type Rd struct{
	        str string
		cur int
	}

func(R *Rd) Read(p []byte)(m int, err error){
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


func MyReader(s string) io.Reader{

	r := Rd{s,0}
	return &r
}


func MyReader_test(){
	y := MyReader("This is freaking awesome!!")
	p := make([]byte, 4)
	for{
		n,err := y.Read(p)
		// you should not print n-p pair after if-statement
		// because you get out of the loop early and leave out the final
		// pair
		fmt.Printf("%d bytes read, Buffer: %v\n",n,string(p))
		if err != nil{
			break
	}

	}
}

//func LimitReader(r io.Reader, n int) (newr io.Reader){
//	//p := make([]byte, 1)
//	total := 0
//	func (newr io.Reader) Read(p []byte)(m int, err error){
//		m,err = r.Read(p)
//		total += m
//		if total > n {
//			m = m-(total-n)
//			err = io.EOF
//			return
//		}
//		return
//
//
//	}
//	return
//
//
//}



//func LimitReader_test(){
//	r := strings.NewReader("woshiyigefeichangliaobuqideren")
//	newr := LimitReader(r,7)
//	p := make([]byte, 3)
//
//	for {
//		nbytes, err := newr.Read(p)
//		if err!= nil{
//			break
//		}
//		fmt.Printf("%v bytes read, Buffer:%v\n",nbytes,string(p))
//
//	}
//
//
//}

func main() {

	//test_Custom_ioReader()
	//test_ioReader()
	MyReader_test()
	//LimitReader_test()

}
