package main
import ("fmt"
	"strings"
)

func ForScopeWithOneVariable() {
	var ops1 []func()
	names := []string{"sfzhang", "xlyang", "lyzhang"}
	for _, name := range names {
		ops1 = append(ops1, func() {
			fmt.Printf("address: %p\n", &name)
			fmt.Println(name)
		})

	}

	for _, print := range ops1 {

		print()

	}
}

func ForScopeWithInnerVariable() {

	var ops2 []func()
	names := []string{"sfzhang","xlyang","lyzhang"}

	for _, name := range names {
		nm := name
		ops2 = append(ops2, func() {
			fmt.Printf("address: %p\n", &nm)
			fmt.Println(nm)
		})
	}

	for _, print := range ops2 {

		print()

	}

}

func ReturnFunc() func(){

	x := 3
	y := func() {
		//fmt.Printf("x{value:%v, address:%p}",operand,&operand)
		//fmt.Printf("Address: %p\n",operand)
		x = x*10
		fmt.Printf("InnerX--->{value:%v, address:%p}\n",x,&x)
	}
	fmt.Printf("OuterX--->{value:%v, address:%p}\n",x,&x)
	return y


}

//func ReferenceType(lst []string, dict map[string]string, ch chan int, f func(), num int) {
func ReferenceType1(intf ...interface{}) {

	for _,j :=range intf {
		d := j
		// grasp the difference between the following two lines of code
		// is essential to fully understand  what happens inside ForScopeWithOneVariable
		//fmt.Printf("Inner Address:%p\t",j)
		//fmt.Printf("Inner Address:%p\t",&j)
		fmt.Printf("Inner Address:%p\t",d)
	}
	fmt.Println()
}

func ReferenceType2(sl []int, mp map[string]string, c chan int, f func(), num int) {

	fmt.Printf("Inner Address:%p\t",sl)
	fmt.Printf("Inner Address:%p\t",mp)
	fmt.Printf("Inner Address:%p\t",c)
	fmt.Printf("Inner Address:%p\t",f)
	fmt.Printf("Inner Address:%p\n",&num)



}
func ReferenceType_Test(){
	lst := []int{1,3}
	dict := make(map[string]string)
	dict["name"]="sfzhang"
	dict["address"]="55 Oakmount Road"
	ch := make(chan int)
	y := func(){fmt.Printf("Printing Something!")}

	nm := 3


	fmt.Printf("Outer Address:%p\t",lst)
	fmt.Printf("Outer Address:%p\t",dict)
	fmt.Printf("Outer Address:%p\t",ch)
	fmt.Printf("Outer Address:%p\t",y)
	fmt.Printf("Outer Address:%p\n",&nm)
        // the following two will get the same results
	// because all the parameters are reference types
	// except the last one, go functions pass by values
	// the values of these types are  with a
	// pointer pointing to the underlying structures
	ReferenceType1(lst, dict, ch, y,nm)
	ReferenceType2(lst, dict, ch, y,nm)
}



func main() {

	ForScopeWithOneVariable()
	fmt.Println(strings.Repeat("*", 180))
	ForScopeWithInnerVariable()

	fmt.Println(strings.Repeat("*", 180))
	ReturnFunc()
	ReturnFunc()
	ReturnFunc()
	ReturnFunc()

	f := ReturnFunc()
	f()
	f()
	f()
	f()

	fmt.Println(strings.Repeat("*", 180))
	ReferenceType_Test()



}
