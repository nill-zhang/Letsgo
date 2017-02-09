package main

import (
	"net"
	"fmt"
	"bufio"
	"io"
)

func test_dial(){

	conn, err := net.Dial("tcp","www.python.org:443")
	if err != nil{panic(err)}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	b := make([]byte, 10)

	for{
		n, err := bufio.NewReader(conn).Read(b)
		fmt.Println("status: ", string(b[:n]))
		if err != nil{
			break
		}
	}
}


func test_listen(){
	l, err := net.Listen("tcp", ":80")
	defer l.Close()
	if err != nil{panic(err)}
	for {
		// wait for a connection
		c, err := l.Accept()
		if err != nil{panic(err)}
		// handle a connection in a go routine
		go func(con net.Conn){

			io.Copy(con, con)
			con.Close()

		}(c)
	}



}

func test_url(){


}


func test_mail(){


}

func main() {
	test_dial()

}

