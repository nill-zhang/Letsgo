package main

import ("fmt";"time")

func main(){
    
    num := 100
    
    switch num{

        case  0    : fmt.Println("num is 0")
        case 10    : fmt.Println("num is 10")
        case 1000  : fmt.Println("num is 1000")
        default    : fmt.Println("default")

    }


    t := time.Now()
    fmt.Println(t)
    switch t.Weekday(){

        case time.Sunday,time.Saturday: fmt.Println("Horray! It's weekend!")
        default:                        fmt.Println("Oh, no, It's weekday!")


    } 




    //switch without expression behaves like if-else
    //and you can omit default as well
    switch{

        case num % 2 == 0: fmt.Println(num , "is an even number")
        case num % 2 != 0: fmt.Println(num , "is not an even number")

    }


}
