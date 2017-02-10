package main

import(
	"strconv"
	"fmt"
	"strings"
)


func  test_strconv_format(){

	bool_str := strconv.FormatBool(true)
	// 102 is byte("f")
	float_str := strconv.FormatFloat(3.1415926,102,4,64)
	hex_str := strconv.FormatInt(-15*256-7*16-7, 16)
	oct_str := strconv.FormatUint(5*64+7*8+2, 8)
	bin_str := strconv.FormatInt(-192, 2)

        p := fmt.Println

	p("bool_str:", bool_str)

	p("float_str:", float_str)
	p("hex_str:", hex_str)
	p("oct_str:", oct_str)
	p("bin_str:", bin_str)

}


func test_strconv_parse(){

	output_fmt :="type: %T	value: %v\n"
	p := fmt.Printf

	b,_ := strconv.ParseBool("true")
	p(output_fmt, b,b)

	f64,_ := strconv.ParseFloat("223.23113", 64)
	p(output_fmt, f64,f64)

	f32,_ := strconv.ParseFloat("223.23113", 32)
	p(output_fmt, f32, f32)

	i2,_ := strconv.ParseInt("11110000", 2, 32)
	p(output_fmt, i2, i2)

	i8,_ := strconv.ParseInt("10246370", 8, 32)
	p(output_fmt, i8, i8)

	i16,_ := strconv.ParseInt("aef01", 16, 32)
	p(output_fmt, i16, i16)

}

func  test_string_format(){

	//alex := Resident{"张少锋",29,70.21}
	p := fmt.Printf
	byte_slice1 := []byte{122,74,75, 93, 109}
	string_slice1 := []string{"h","e","l", "l", "o"}
	string_slice2:= []string{"sfe", "你好","！~", "✿", "帅气"}
        name := "少锋张"

	p("%s,%s,%s,%s\n",byte_slice1, string_slice1, string_slice2, name)
	p("%q,%q,%q,%q\n",byte_slice1, string_slice1, string_slice2, name)
	p("%x,%x,%x,%x\n",byte_slice1, string_slice1, string_slice2, name)
	p("%X,%X,%X,%X\n",byte_slice1, string_slice1, string_slice2, name)
	p("%v,%v,%v,%v\n",byte_slice1, string_slice1, string_slice2, name)
	p("%#v,%#v,%#v,%#v\n",byte_slice1, string_slice1, string_slice2, name)
        // note that if you use 30/80, you will get 0, because the result is int
	p("%4.2f\n",30.0/80.0*100)
	// %% is for %
	p("%4.2f%%\n",30.0/80.0*100)
	p("%t\n", 3==3)
	p("[Binary: %b][Octal: %o][Hex: %x][Hex: %X]\n",192,192,192,192)
	p("[Binary: %#b][Octal: %#o][Hex: %#x][Hex: %#X]\n",192,192,192,192)
	fmt.Println(strings.Repeat(fmt.Sprintf("%c",0x2618),15))
	p("%U\n",0x2618)
	p("%U\n",9752)
	p("%#U\n",0x2618)
	p("%#b\n",0x2618)
	p("%[2]s	%[1]s\n", "是张少锋","我的名字")

	characters_hex := []rune{0x767d, 0x9d6c, 0x7fd4}
	p("%s\n", string(characters_hex))
	characters_decimal := []rune{30333, 40300, 32724}
	p("%s\n", string(characters_decimal))

	// characters_u := []rune{\u767d, \u9d6c, \u7fd4}
	// p("%s\n", string(characters_u))

	// it must be single quote
	characters_x := []rune{'\x76', '\x7d', '\x9d', '\x6c', '\x7f', '\xd4'}
	p("%s\n", string(characters_x))

	p(fmt.Sprintf("%#X\n", "张"))

	myname := string("张少锋")
	b := []rune(myname)
	p("%U\n",b)
	p("%#U\n",b)
	p("%c",b)




}

func main() {
	//test_strconv_format()
	//test_strconv_parse()
	test_string_format()


}

