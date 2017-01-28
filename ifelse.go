package main

import "fmt"

//function main must have no arguments and no return values
func main(){

    if num := 10; num < 0{

        fmt.Println("num is negative")

    }else if num < 10{

        fmt.Println("num has one digit")


    }else{

        fmt.Println("num has more than one digit")

    }


    if 8%2 == 0{

        fmt.Println("8 is an even number")
    }



}
