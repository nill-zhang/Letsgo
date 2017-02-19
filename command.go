package main

import ("flag"
        "fmt"
	"strings"
)


func main() {
	var n = flag.Bool("n", false, "omit trailing newline")
	var s = flag.String("s", " ", "separator between arguments")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*s))
	if !*n{
		fmt.Println()
	}
}
