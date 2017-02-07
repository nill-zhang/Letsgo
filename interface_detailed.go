package main
import  "fmt"

type Person struct{
	name string
	age uint8
	weight float64

}

type  Intslice []int
type  Floatslice []float64
type  Personslice []Person

type Collection interface{

	len() int
	get(int) interface{}
	bigger(int, int) bool
}

// implement len()
func (i Intslice) len() int {return len(i)}
func (f Floatslice) len() int {return len(f)}
func (p Personslice) len() int {return len(p)}

// implement get()
func (i Intslice) get(idx int) interface{}{return i[idx]}
func (f Floatslice) get(idx int) interface{}{return f[idx]}
func (p Personslice) get(idx int) interface{}{return p[idx]}

// implement bigger(), note here you can not use f.get(m)>f.get(n)
// because both f.get return any type(interface {}) which does not
// implement > operation
func (i Intslice) bigger(m,n int) bool {return i[m] > i[n]}
func (f Floatslice) bigger(m,n int) bool {return f[m] > f[n]}
func (p Personslice) bigger(m,n int) bool {return p[m].age > p[n].age}

func Getmax(c Collection) (maxelement interface{}){

	if c.len() == 0{return nil}
	maxelement = c.get(0)
	maxindex := 0
	for i:=1; i<c.len(); i++{

		if c.bigger(i, maxindex){
			maxelement = c.get(i)
			maxindex = i
		}
	}
	return maxelement

}


func test_Getmax(){

	// test_data is a slice of type Collection, because Intslice, Floatslice, Personslice
	// all implement Colleciton interface, so they can be a Collection type variable
	// function Getmax requires a collection type, whatever that is , as long as it
	// implements its mandatory methods(len(), get(), bigger())
	var test_data = []Collection{Intslice{2,3,4,2,7,8,4,2,90,23,782,51,1,59,41,0,-13,-4,9},
		Floatslice{2.11,-232, -21.012, 22.4, 87.21, 18.22, 83.09, 901.12, 8.1},
		Personslice{{"Alex",23,78},{"John", 13, 55}, {"Helen",31,60},
			    {"Nova",21,66},{"Michael",36, 80.00},{"Sharlane",25,68}}}

	for _,value :=range test_data {

		fmt.Printf("Slices: %v	Max: %v\n", value, Getmax(value))

	}

}
func main(){

	test_Getmax()

}
