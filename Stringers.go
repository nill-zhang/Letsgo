package main

import (
	"fmt"
	"strings"

)

type Ipaddress [4]byte

type Person struct {
	name string
	age uint8
	weight float32

}

func (i Ipaddress)String() string{
	return  fmt.Sprintf("%v.%v.%v.%v\n",i[0],i[1],i[2],i[3])
}

func (p Person)String() string{
	a := strings.Repeat("*",60)
	return  a + fmt.Sprintf("\nName: %v-----Age: %v years-old" +
		"-----Weight: %v KG\n",p.name, p.age,
		p.weight)+ a

}

func main() {

	ips := map[string]Ipaddress{
		"google dns" : {224,222,44,88},
		"loop back" : {127,0,0,1},
	}

	for name,ip := range ips{

		fmt.Printf("\033[35m[Host]\033[0m: %v \033[34m[Ip]\033[0m: %v", name, ip)
	}

	myprofile := Person{age : 29, name: "Shaofeng Zhang", weight: 71,}
	fmt.Println( myprofile)


}

// conclusion:
// [1].
// Stringer is an interface, when a concrete type implements its String()
// method, when you use fmt to print the value, fmt will look for its string() method
//
// [2].
// using + to break string to multiple lines
//
// [3].
// if struct values are in one line the last comma can be omitted, for example
// a:= Person{age:24,name:"Alex",weight:71}
// instead of a := Person{age : 24,
//                        name: "Alex",
//                        weight: 71,}