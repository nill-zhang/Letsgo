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

func toMulticastAddress(address flag) flag { return address | multicastAddressFlag}
func flipLoopbackAddress(address  flag) flag{return address ^ loopbackAddressFlag}
func toNotAnycastAddress(address  flag) flag{return address &^ anycastAddressFlag}
func isCastAddress(address flag) bool {
	return (address & (broadcastAddressFlag | multicastAddressFlag | anycastAddressFlag)) != 0
}

func main() {

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
