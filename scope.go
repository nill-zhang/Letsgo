package main
import ("fmt"
)


func main() {

	var prints []func()
	names := []string{"sfzhang","xlyang","lyzhang"}
	for _, name := range names{
		prints = append(prints, func(){fmt.Println(name)})

	}

	for _, print := range prints{

		print()

	}


}
