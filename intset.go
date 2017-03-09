package main

import (
	"bytes"
	"strconv"
	"fmt"
)

type IntSet struct{
	set []uint64
}

// len returns the length of the set, how many elements are in the set
func (s *IntSet)Len() uint64{

	return uint64(len(s.Elem()))

}

func (s *IntSet)UnionWith(snd IntSet){

	if len(s.set) <= len(snd.set){
		for i:=0;i<len(s.set);i++{
			s.set[i] = s.set[i] | snd.set[i]

		}
		s.set = append(s.set, snd.set[len(s.set):]...)
	}

	for i:=0;i<len(snd.set);i++{
			s.set[i] = s.set[i] | snd.set[i]

		}




}

func (*IntSet)IntersectWith(){


}

func (*IntSet)DifferenceWith(){


}

func (*IntSet)AddAll(sets ...IntSet){


}


func (s *IntSet)Remove(nm uint64) bool{
        if s.Has(nm){

		s.set[nm/64] = s.set[nm/64] &^ (uint64(1)<<(nm%64))
		return true
	}


	return false
}

func (s *IntSet)Clear(){

	for i:=0;i<len(s.set);i++{
		s.set[i] = s.set[i] & uint64(0)
	}



}


func (*IntSet)Copy(){






}

func (s *IntSet)Elem() (res []uint64){

	for i:=0;i<len(s.set);i++{
		for j:=0;j<64;j++{
			if (uint64(1)<<uint64(j) & s.set[i]) != 0{
				res = append(res, uint64(i*64+j))

			}

		}

	}

	return

}


func (s *IntSet)Add(nu uint64){

	//fmt.Printf("%d\n",s.Len())
	//fmt.Printf("%d\n",len(s.set))
	//fmt.Printf("%d\n",nu/64)

	if nu/64 >= uint64(len(s.set)){
		y := make([]uint64, (nu/64)+1-uint64(len(s.set)))
		s.set = append(s.set, y...)
		//fmt.Printf("len(s.set):    %d\n",len(s.set))


	}
	s.set[nu/64] = s.set[nu/64] | (uint64(1)<<(nu%64))
	//fmt.Printf("s.set[%v]:    %d\n",nu/64, s.set[nu/64])


}


func (s *IntSet)Has(nu uint64) bool{

	if nu/64 > uint64(len(s.set)-1){
		return false
	}
	if s.set[nu/64] & (uint64(1) << (nu%64)) != 0{
		return true
	}
	return false


}


func (s *IntSet)SymmetricDifference(){



}

func (s *IntSet)String() string{

	var buf bytes.Buffer
	_, err := buf.Write([]byte("{"))
	//fmt.Printf("s.Elem(): %v\n", s.Elem())
	for _,j := range s.Elem(){

		_, err = buf.Write([]byte(strconv.FormatUint(j,10)))
		_, err = buf.Write([]byte(","))
	}

	_, err = buf.Write([]byte("}"))

	if err != nil{

		 return fmt.Sprintf("%v while writing to buffer\n", err)
	}
	return string(buf.Bytes())

}

func main(){
	p := fmt.Printf
	//var myset *IntSet
	myset := &IntSet{}
	myset.Add(uint64(6))
	myset.Add(uint64(66))
	myset.Add(uint64(666))
	myset.Add(uint64(6666))
	p("myset: %v\n", myset)
	p("myset has 65? %t\n", myset.Has(uint64(65)))
	p("myset has 66? %t\n", myset.Has(uint64(66)))
	p("myset has 666 %t\n", myset.Has(uint64(666)))
	p("myset has 665? %t\n", myset.Has(uint64(665)))


	var herset =&IntSet{}

	herset.Add(uint64(1))
	herset.Add(uint64(10))
	herset.Add(uint64(100))
	herset.Add(uint64(16))
	herset.Add(uint64(6666))
	p("herset: %v\n", herset)
	p("myset: %v\n", myset)
	p("Lenth of myset: %v\n", myset.Len())

	myset.UnionWith(*herset)
	p("After UnionWith() myset: %v\n", myset)

	ok := myset.Remove(uint64(100))
	if ok{
		p("After Remove(100) myset: %v\n", myset)
	}

	myset.Clear()
	p("After Clear() myset: %v\n", myset)






}