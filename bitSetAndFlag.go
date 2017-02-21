package main

import "fmt"

type flag uint8

const (
	broadcastAddressFlag  = 1 << iota // 00000001
	multicastAddressFlag         // 1<<1 00000010
	loopbackAddressFlag          // 1<<2 00000100
	anycastAddressFlag           // 1<<3 00001000
)
func isBroadcastAddress(address flag) bool {
	return (address & broadcastAddressFlag) == broadcastAddressFlag
}
func isMulticasttAddress(address flag) bool {
	return   (address & multicastAddressFlag) == multicastAddressFlag
}
func isLoopbackAddress(address flag) bool {
	return (address & loopbackAddressFlag) == loopbackAddressFlag
}
func isAnycastAddress(address flag) bool {
	return (address & anycastAddressFlag) == anycastAddressFlag
}


// Note that we didn't modify the input, we return a new one, if you
// want to modify the input use its pointer as an argument and do the
// *address |= multicastAddressFlag alike
func toMulticastAddress(address flag) flag { return address | multicastAddressFlag}
func flipLoopbackAddress(address  flag) flag{return address ^ loopbackAddressFlag}
func toNotAnycastAddress(address  flag) flag{return address &^ anycastAddressFlag}
func isCastAddress(address flag) bool {
	return (address & (broadcastAddressFlag | multicastAddressFlag | anycastAddressFlag)) != 0
}

func address_test(){
	var noBroadcast, broadcast flag = 240, 17
	var noMulticast, multicast flag = 13, 34
	var noLoopback, loopback flag = 11, 69
	var noAnycast, anycast flag = 7, 31
        p := fmt.Printf
	p("%08b is broadcast address? [%t]\n",noBroadcast, isBroadcastAddress(noBroadcast))
	p("%08b is  broadcast address? [%t]\n",broadcast , isBroadcastAddress(broadcast))
	p("%08b is multicast address? [%t]\n", toMulticastAddress(noMulticast),
		                           isMulticasttAddress(toMulticastAddress(noMulticast)))
	p("%08b is multicast address? [%t]\n", toMulticastAddress(multicast),
		                           isMulticasttAddress(toMulticastAddress(multicast)))
	p("%08b is loopback address? [%t]\nafter conversion it is %08b, loopback address? [%t]\n",
		 noLoopback,isLoopbackAddress(noLoopback), flipLoopbackAddress(noLoopback),
	isLoopbackAddress(flipLoopbackAddress(noLoopback)))
	p("%08b is loopback address? [%t]\nafter conversion it is %08b, loopback address? [%t]\n",
		 loopback,isLoopbackAddress(loopback), flipLoopbackAddress(loopback),
	isLoopbackAddress(flipLoopbackAddress(loopback)))

	p("%08b is anycast address? [%t]\n", toNotAnycastAddress(noAnycast),isAnycastAddress(toNotAnycastAddress(noAnycast)))
	p("%08b is anycast address? [%t]\n", toNotAnycastAddress(anycast),isAnycastAddress(toNotAnycastAddress(anycast)))

	p("%08b is cast address? [%t]\n", loopback, isCastAddress(loopback))
	p("%08b is cast address? [%t]\n", noLoopback, isCastAddress(noLoopback))
	p("%08b is cast address? [%t]\n", anycast, isCastAddress(anycast))
	p("%08b is cast address? [%t]\n", noMulticast, isCastAddress(noMulticast))
}

func set_test(){
	var x, y uint8 = 1<<4|1<<3|1<<7, 1<<2|1<<5|1<<3
        fmt.Printf("x: %08b, Y: %08b\n", x,y)
	fmt.Printf("x|y:  %08b\n",x|y)  // union
	fmt.Printf("x&y:  %08b\n",x&y)  // intersection
	fmt.Printf("x^y:  %08b\n",x^y)  // symmetric difference
	fmt.Printf("x&^y: %08b\n",x&^y) // x-y
	fmt.Printf("y&^x: %08b\n",y&^x) // y-x
	// test membership
	for i:=uint8(0);i<=7;i++{
		if x&(1<<i) != 0{
			fmt.Printf("%d bit in x is on\n",i)
		}

	}



}

func main() {
	set_test()



}
