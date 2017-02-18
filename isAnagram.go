package main

import ("fmt"
)

func isMapEqual(a, b map[rune]int) bool{

	if len(a) != len(b){
		return false
	}
	for i := range a{
		if a[i] != b[i]{
			return false
			break
		}
	}

	for j := range b{
		if b[j] != a[j]{
			return false
		        break}
	}

	return true
}

func isAnagram(a,b string) bool{
	counta := make(map[rune]int)
	countb := make(map[rune]int)

	// when we use range for loop on a string
	// we iterate over its runes
	// see byteOrRune function for detail
	for _,j := range a{
		counta[j]++
	}

	for _,j := range b{
		countb[j]++
	}

	// we can not compare two map type directly using counta == countb
	// we can not compare slices directly as well
	return isMapEqual(counta, countb)



}

func isAnagram_test(teststr1, teststr2 string){

	fmt.Printf("%v is a anagram of %v : " +
		"%t\n",teststr1, teststr2, isAnagram(teststr1, teststr2))


}

func byteOrRune(){

	a := "我是中国人"
	b := "sfe"
	for _,j := range a{
		fmt.Printf("a(range loop): [type of elements: %T]\t[value of elements: %[1]v]\n", j)
	}

	for i:=0;i<len(a);i++{
		fmt.Printf("a(iter  loop): [type of elements: %T]\t[value of elements: %[1]v]\n", a[i])
	}

	for _,j := range b{
		fmt.Printf("b(range loop): [type of elements: %T]\t[value of elements: %[1]v]\n", j)
	}

	for i:=0;i<len(b);i++{
		fmt.Printf("b(iter  loop): [type of elements: %T]\t[value of elements: %[1]v]\n", b[i])
	}




}

func main() {
	isAnagram_test("我是王八蛋但","是我但蛋王八")
	// should be false because ？? is different
	isAnagram_test("beini搞败了？","?搞败了nbiie")
	isAnagram_test("sfjelfjeop!@#$%^", "sfjelfje$%op!@#^")
	isAnagram_test("1234567", "7654321")

}
