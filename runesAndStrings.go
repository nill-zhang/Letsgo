package main
import ("fmt"
	"unicode/utf8"

)

var p = fmt.Printf
func runesInString(){
	address1 := "安徽省凤凰城家家景园商住3单元605室"
	address2 := "404- 55 Oakmound Road"
	address3 := "合肥市蜀山区蜀南庭苑310"
	address4 := "ギャンブル等依存症対策基本法案を掲載しました"
	addresses := []string{address1,address2,address3,address4}

	for _, add := range addresses{
	      	var n,m int32
	     	 //for range add{
			//  n++
	     	 //}
		p("String: %v has: ", add)
		//for i:=0;i<len(add);i++{
		for i:=0;i<len(add);{
			r, s := utf8.DecodeRuneInString(add[i:])
			p("%c, ",r)
			i += s
			n++

		}
		p("%v runes!\n",n)

		// for loop actually iterate over runes
		p("String: %v has: ", add)
		for i,j := range add{
			m++
			p("i:%v,%c, ",i,j)
		}
		p("%v runes!\n",m)
	}




}

func rune_construction(){
	string_encoded_utf8 := "\xe5\xae\x89\xe5\xbe\xbd\xe7\x9c\x81\xe5" +
			"\x87\xa4\xe5\x87\xb0\xe5\x9f\x8e\xe5\xae\xb6" +
			"\xe5\xae\xb6\xe6\x99\xaf\xe5\x9b\xad\xe5\x95" +
			"\x86\xe4\xbd\x8f3\xe5\x8d\x95\xe5\x85\x83605\xe5\xae\xa4"
	string_encoded_gb2312 := "\xbd\xe1\xbb\xe9\xd2\xaa\xd5\xd2\xb0\xa1\xb8\xf6"
	string_literal := "我是中国人"
	// some times strings are hard to type or invisible, so we can use \u for 16bit unicode
	// \U for 32bit unicode code points to represent those underlying character
	string_u := "\u4e00\u4e11\u4e27\u4e58\u4eb3"
	string_U := "\U00004e16\U0000754c"
	runes_hex := []rune{0x8d3e, 0x77eb,0x8d70,0x8d73}
	runes_decimal := []rune{3443,25830,37089,23569}
	// runes is of type int32, why we can string to represent them??
	runes_u := []rune{'\u767d', '\u9d6c', '\u7fd4','\U00004e16'}
	// rune literal must be in single quote
	// a rune whose value is less than 256 can be excaped using \x
	// for high values \u,\U must be used to escape
	// '\xe4\xb8\x96' is not a valid rune this is a code point encoded
	// as a string
	runes_x := []rune{'\x76', '\x7d', '\x9d', '\x6c', '\x7f', '\xd4'}
	p("runes_hex:%c\n", runes_hex)
	p("runes_decimal:%c\n", runes_decimal)
	p("runes_u:%c\n", runes_u)
	p("runes_x:%c\n", runes_x)

	// note that %c/%u/%U only works for rune rune[]
	// in accordance with code point
	p("string(runes_hex):%s\n", string(runes_hex))
	p("string(runes_decimal):%s\n", string(runes_decimal))
	p("string(runes_u):%v\n", string(runes_u))
	p("string(runes_x):%v\n", string(runes_x))


	p("string_encoded_utf8:%s\n", string_encoded_utf8)
	// the default decode is utf-8, so this will print nonsense
	p("string_encoded_gb2312:%s\n", string_encoded_gb2312)
	p("string_literal:%s\n", string_literal)
	p("string_u:%s\n", string_u)
	p("string_U:%s\n", string_U)

	// 张少锋
	name := "\xe5\xbc\xa0\xe5\xb0\x91\xe9\x94\x8b"
	p("name: %v\n", name)
        // Note the difference between the following two
	// %x can be used on both string and rune/rune[], it will display
	// the utf-8 encoded bytestring for string  and will display unicode
	// code point for rune/rune[]
	p("name: \033[36m%# x\033[0m\n", name)
	p("name: %#x\n", name)
	rune_name := []rune(name)
	p("rune_name: %#x\n", rune_name)


}


func main(){
 	rune_construction()
	runesInString()
}

