package main

import (
	"fmt"
	"os"
	"text/template"
)

type Inventory struct{
	material string
	count uint

}


func test_templatefile(){
	tpl, err := template.ParseFiles("letter.php")
	if err != nil {
		fmt.Println("There was an error parsing file", err)
	}

	friends := []string{"Alex", "Conor", "Ken", "Ronnie", "Patick", "Nina", "Jeremy", "Gentry", "Christian"}

	err = tpl.Execute(os.Stderr, friends)
	if err != nil {
		fmt.Println("There was an error executing template", err)
	}
}

func test_template_string(){

	wines := Inventory{"Yellow tail", 500}
	tpl,err := template.New("test").Parse("Currently we have {{.count} {{.material}} bottle wines")
	if err != nil{
		fmt.Println("An error has happened while parsing template string")
	}
	tpl.Execute(os.Stdout,wines)
	if err != nil{panic(err)}



}
func main() {

	test_template_string()
}
