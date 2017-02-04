package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{

	return  fmt.Sprintf("Can not Sqrt negative number: %v", float64(e))


}

func Sqrt(i float64) (float64, error){

	if i<0 {
		return 0,ErrNegativeSqrt(i)
	} else{

		return  math.Sqrt(i),nil
	}


}

func main() {

	var  num =[]float64{-32.32, 34.333}

	for _, n := range num{

		if result,err:= Sqrt(n); err!=nil{

			fmt.Println(err)
		}else{

			fmt.Println("sqrt result:",result)
		}

	}


}
