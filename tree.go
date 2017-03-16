package main

import (
	"errors"
	"fmt"
)

type BinarySearchTree struct{

	value int
        lchild,rchild,parent,sibling *BinarySearchTree
}

// InitFromSlice takes a slice and add all the elements to the tree
// the first element will be the root node
func (b *BinarySearchTree)InitFromSlice(s []int)error {
	if len(s) < 1 {
		return errors.New("no elements to initialize from")
	}
	b.value = s[0]
	for _, item := range s[1:] {
		b.Add(item)
	}
	return nil
}

// Sort will output the tree as a sorted slice
// breadth-firt leftchild-->value-->rightchild
func (b *BinarySearchTree)Sort()(sorted []int){
	if b.lchild != nil {
		sorted = append(sorted,b.lchild.Sort()...)
	}
	sorted = append(sorted, b.value)
	if b.rchild != nil {

		sorted = append(sorted,b.rchild.Sort()...)

	}

	return

}

func (b *BinarySearchTree)Add(elemvalue int){
		if b.value > elemvalue {
			if b.lchild != nil {
				b.lchild.Add(elemvalue)
				return
			}
                        // if no left child present, then construct a left child and
			// set its value as elemvalue, its parent as b
			newnode := &BinarySearchTree{value:elemvalue,parent:b, sibling:b.rchild}
			b.lchild = newnode
			return
		}
		if b.rchild != nil {
			b.rchild.Add(elemvalue)
			return
		}
		newnode := &BinarySearchTree{value:elemvalue,parent:b,sibling:b.lchild}
		b.rchild = newnode
}

func (b *BinarySearchTree)AddAll(elems []int){

	for _,j := range elems{
		b.Add(j)
	}


}

// Has checks whether the tree has an element node whose value matches the given parameter
// depth-first node-->lchild--->rchild
func (b *BinarySearchTree)Has(num int)bool{

	if b.value == num{
		return true
	}
	// check left child
	if b.lchild != nil && b.lchild.Has(num) {
		return true
	}
	// check right child
	if b.rchild != nil && b.rchild.Has(num) {
         	return true
        }
         return false


}


func (b *BinarySearchTree)Max() int{

	return
}


func (b *BinarySearchTree)Min() int{

	return
}

func (b *BinarySearchTree)Height() int{

	return
}


func BinarySearchTreeInitial_test(){

	bt := &BinarySearchTree{34,nil,nil,nil,nil}
	bt.InitFromSlice([]int{-21,-1,-5,-8,-17,20,3,9,16,13,2,-76,788,28,-22})
	fmt.Printf("sorted: %v\n",bt.Sort())
	bt.AddAll([]int{56,78,-22,-46,304})
	fmt.Printf("sorted: %v\n",bt.Sort())
	fmt.Printf("bt has 56? %t\n",bt.Has(56))
	fmt.Printf("bt has 78? %t\n",bt.Has(78))
	fmt.Printf("bt has 65? %t\n",bt.Has(65))
	fmt.Println(bt)


 }



func main() {
	BinarySearchTreeInitial_test()
}
