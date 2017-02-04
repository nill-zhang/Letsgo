package main
import "fmt"


type NameError struct{
	errcode uint16
	name  string

}

func (n *NameError) Error() string{

        return  fmt.Sprintf("Error code: %v, can not find %v",
		            n.errcode, n.name)
}


func GetPasswd(name string, passwds map[string]string) error{

	passwd, ok := passwds[name]
	if ok == false {
		return  &NameError{404, name}
	} else{
		fmt.Printf("%v's password is %v\n", name, passwd)
		return nil
	}

}

func main() {

	shadow := map[string]string{"John" : "zsfefe!$#%",
		                    "Alice" : "fefjeo98u902i",
		                    "Tom" : "sfefoj20%^&*( "}

	names_to_search := []string{"Bob","Alice","Sam"}
	for _, n := range names_to_search{

		if err := GetPasswd(n,shadow); err != nil {
			fmt.Println(err)
		}
	}



}

// Behind The Scenes:
// functions often return an error when something wrong with its call
// this error is an interface value
// type error interface {
//    Error() string
// }
// which was assigned a concrete type(NameError) and value in the function
// when fmt print this error it will call the underlying concrete type's Error() method