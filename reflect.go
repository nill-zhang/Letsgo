package main

import (
	"fmt"
	"reflect"
)
type LinkedList struct{

	Value int
        Next *LinkedList
}

// Reflect reconstruct the struct definition and variable assignment
func Reflect(src interface{}){
	st := reflect.ValueOf(src).Elem()
	//typ := reflect.TypeOf(src)== st.Type()
	//mt.Printf("Typ: %v\n",typ)

	Def := ""
	Ass := ""
	Def += fmt.Sprintf("Type %v struct {\n",st.Type())
	Ass += fmt.Sprintf("var l *%v = &%[1]v{",st.Type() )
	for i:=0;i<st.NumField();i++{
		typeField := st.Type().Field(i)
		valueField := st.Field(i)
		Def += fmt.Sprintf("    %s  %v\n",typeField.Name,typeField.Type)
		Ass += fmt.Sprintf("%v:%v,",valueField.Type(),valueField.Interface())


	}

	Def += fmt.Sprintf("}")
	Ass += fmt.Sprintf("}")
	fmt.Println(Def)
	fmt.Println(Ass)




}

// Copy returns shallow copy of src
func Copy(src []*LinkedList) *LinkedList{

	dst := []*LinkedList{new(LinkedList)}
	// reflect.copy takes reflect.Value argument
	// and Their Kind() must be slice or array with the
	// same element type
	num := reflect.Copy(reflect.ValueOf(dst), reflect.ValueOf(src))
	fmt.Printf("Number of Elements Copied: %v\n",num)
	fmt.Printf("Dst:%v\n",dst[0])
	return dst[0]

}

func main() {
     	src := &LinkedList{111,&LinkedList{222,&LinkedList{333,&LinkedList{444,nil}}}}
	Reflect(src)
	cpy := Copy([]*LinkedList{src})
	fmt.Printf("Dst.Value:%v, Dst.Next:%v\n",cpy.Value,cpy.Next)

}
