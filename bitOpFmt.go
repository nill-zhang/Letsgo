package main

import (
	"fmt"
	"strings"
)


func Bitop(lo int64, ro int64, ops []string){

        fpf := fmt.Printf
        entryPrint := func(op string, result interface{}){
		// - means left alignment, otherwise it will be right aligned
		// %032b means total 32 bit, pad with 0
		fpf(strings.Repeat("*",80)+"\n")
		fpf("%-18s: Decimal: %8d Binary: %032b\n","left operand", lo, lo)
		fpf("%-18s: Decimal: %8d Binary: %032b\n","right operand", ro, ro)
		fpf("%-18s: Decimal: %8d Binary: %032b\n", "Bit "+op+" Operation",result,result)

        }

        for _,j := range ops{

		switch j{
			case "&": entryPrint(j, lo&ro)
			case "|": entryPrint(j, lo|ro)
			case "^": entryPrint(j, lo^ro)
			case "+": entryPrint(j, lo+ro)
			case "-": entryPrint(j, lo-ro)
			// bit shift require shift count to be unsigned integer
			case ">>": entryPrint(j, lo>>uint64(ro))
			case "<<": entryPrint(j, lo<<uint64(ro))

			default: fpf("not a valid operator %v", j)
			}

		}
}



func main() {

     	Bitop(6,10,[]string{">>","<<", "&", "|", "^"} )
	Bitop(6,-10,[]string{">>","<<", "&", "|", "^"} )
	Bitop(-6,-10,[]string{">>","<<", "&", "|", "^"} )
	Bitop(-6,10,[]string{">>","<<", "&", "|", "^"} )
}

