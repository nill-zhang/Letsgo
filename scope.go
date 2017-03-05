package main
import ("fmt"
	"strings"
	"os"
)

func ForScopeWithOneVariable() {
	/////////////////////////////////////////////////////////////////
	var ops1 []func()
	names := []string{"sfzhang", "xlyang", "lyzhang"}
	for _, name := range names {
		ops1 = append(ops1, func() {
			fmt.Printf("address: %p\n", &name)
			fmt.Println(name)
		})

	}

	// when you start executing those prints, They will
	// evaluate variable names and find they all share the
	// same name
	for _, print := range ops1 {

		print()

	}
	//////////////////////////////////////////////////////////////////
	fmt.Println(strings.Repeat("*",180))
	letters := []string{"a","b","c","d"}
	var funcs []func()

	for i:=0;i<len(letters);i++{
		funcs = append(funcs, func(){fmt.Println(i)})
	}
	// the following for loop will cause panic, because when you start executing
	// those prints, they all will evaluate variable i, at that time i == 4
	// due to for loop, letters[4] yields index out of range errors
	// for i:=0;i<len(letters);i++{
	//	 funcs = append(funcs, func(){fmt.Println(letters[i])})
	// }

	for _,print :=range funcs{
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

        ///////////////////////////////////////////////////////////////
	fmt.Println(strings.Repeat("*",180))
	letters := []string{"a","b","c","d"}
	var funcs []func()

	for i:=0;i<len(letters);i++{
		temp := letters[i]
		funcs = append(funcs, func(){fmt.Println(temp)})
	}

	for _,print :=range funcs{
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


func IfScope(){
        var c uint8
	if a:=-3; a>0 {
		c = 8
		fmt.Printf("%v\n",a)
	}else {
		a = -a
		fmt.Printf("%v\n",a)

	}
	// undefined: a
	// fmt.Printf("%v\n",a)
	fmt.Printf("%v\n",c)


}

func ForScope(){

	var name string
	for i:=1;i<3;i++{
		var age string
		fmt.Printf("inside for name: %v, age:%v",name,age)


	}
	// undefined: i
	// undefined: age
	// fmt.Printf("%v",i)
	// fmt.Printf("%v",age)

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
	IfScope()
	ForScope()




}
