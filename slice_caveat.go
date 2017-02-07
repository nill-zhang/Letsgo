package main
import ("fmt";"math/rand")

func modify1(slice *[]int){
    	*slice = (*slice)[:3]
	fmt.Println("within modify1:", *slice)

}


func modify2(slice []int){
        fmt.Printf("%p\n", &(slice[1]))
    	slice = slice [:3]
	fmt.Printf("%p\n", &(slice[1]))
	fmt.Println("within modify2:", slice)

}

func modify3(slice []int){
    	slice[0] = 3333
	fmt.Println("within modify3:", slice)


}

func main(){
    	original1:= []int {1, 2, 3, 4, 5}
   	original2:= []int {1, 2, 3, 4, 5}
   	original3:= []int {1, 2, 3, 4, 5}

  	modify1(&original1)
   	fmt.Println("original1: = ", original1)

   	modify2(original2)
    	fmt.Println("original2: = ", original2)

    	modify3(original3)
    	fmt.Println("original3: = ", original3)

        test := make([]int, 7)
	for i:=0;i<7;i++{
		test[i] = rand.Intn(20)
	}
	fmt.Println("test: = ", test)
	modify2(test)
	fmt.Println("test: = ", test)
}

/*
Result:
within modify1: [1 2 3]
original1: =  [1 2 3]
0x10432224
0x10432224
within modify2: [1 2 3]
original2: =  [1 2 3 4 5]
within modify3: [3333 2 3 4 5]
original3: =  [3333 2 3 4 5]

In my 3 modify functions, the second does not change the original slice,
the other two changes the underlying data, in the second modify
we just using a local variable slice assigned to point to the original slice
it has no effect on the original one

*/
