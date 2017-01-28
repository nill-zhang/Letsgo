package main

import "fmt"

func main(){

    for i:=1;i<=10;i++{

    /*Indentation doesn't really matter in golang*/
    fmt.Println(i*i)

}

    var j int = 10

    for j>0{

        //Note: convert j to its string format, which will surprise you
        //      it gets empty string
        fmt.Println(string(j))
        j--


}


    for{

        fmt.Println("before break out the loop")
        break  
 
}
    //print odd number
    for l:=20;l>=0;l--{
    
        if l%2 ==0{
         
            continue
        }
        

        fmt.Println(l)
}


}
