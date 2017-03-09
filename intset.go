package main

import (
	"bytes"
	"strconv"
	"fmt"
	"text/tabwriter"
	"os"
	"strings"
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
		        // keep the rest s.set[len(snd):]

		}
}


func (s *IntSet)IntersectWith(snd IntSet){
	if len(s.set) <= len(snd.set){
		for i:=0;i<len(s.set);i++{
			s.set[i] = s.set[i] & snd.set[i]

		}
	}

	for i:=0;i<len(snd.set);i++{
			s.set[i] = s.set[i] & snd.set[i]
		        // truncate the set
		        s.set = s.set[:len(snd.set)]

		}
}


func (s *IntSet)DifferenceWith(snd IntSet){

	if len(s.set) <= len(snd.set){
		for i:=0;i<len(s.set);i++{
			s.set[i] = s.set[i] &^ snd.set[i]

		}
	}
	for i:=0;i<len(snd.set);i++{
			s.set[i] = s.set[i] &^ snd.set[i]
		        // keep the rest s.set[len(snd.set):]
		}
}


func (s *IntSet)SymmetricDifference(snd IntSet){
	if len(s.set) <= len(snd.set){
		for i:=0;i<len(s.set);i++{
			s.set[i] = (s.set[i] &^ snd.set[i]) |  (snd.set[i] &^ s.set[i])

		}

		s.set = append(s.set, snd.set[len(s.set):]...)
	}

	for i:=0;i<len(snd.set);i++{
			s.set[i] = (s.set[i] &^ snd.set[i]) |  (snd.set[i] &^ s.set[i])
		        // keep the rest s.set[len(snd.set):]
		}
}


func (s *IntSet)AddAll(sets ...IntSet){
	for _,j := range sets{
		for _,k := range j.Elem(){
			s.Add(k)
		}
	}
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


func (s *IntSet)Copy() *IntSet{
	cp := make([]uint64, len(s.set))
	copy(cp, s.set)
	return &IntSet{cp}

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


	if nu/64 >= uint64(len(s.set)){
		y := make([]uint64, (nu/64)+1-uint64(len(s.set)))
		s.set = append(s.set, y...)

	}
	s.set[nu/64] = s.set[nu/64] | (uint64(1)<<(nu%64))


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




func (s *IntSet)String() string{

	var buf bytes.Buffer
	elements := []string{}
	//_, err := buf.Write([]byte("{"))
	for _,j := range s.Elem(){

		elements = append(elements, strconv.FormatUint(j,10))
		//_, err = buf.Write([]byte(strconv.FormatUint(j,10)))
		//_, err = buf.Write([]byte(","))
	}

	_, err := buf.Write([]byte("{"+strings.Join(elements,",")+"}"))

	if err != nil{

		 return fmt.Sprintf("%v while writing to buffer\n", err)
	}
	return string(buf.Bytes())

}


func SetOperation_test(){
	myset := &IntSet{}
	myset.Add(uint64(6))
	myset.Add(uint64(66))
	myset.Add(uint64(666))
	myset.Add(uint64(6666))
	myset.Add(uint64(3))

	var herset =&IntSet{}
	herset.Add(uint64(6))
	herset.Add(uint64(10))
	herset.Add(uint64(100))
	herset.Add(uint64(18))
	herset.Add(uint64(6666))

	s1 := myset.Copy()
	s2 := myset.Copy()
	s3 := myset.Copy()
	s4 := myset.Copy()

	s1.UnionWith(*herset)
	s2.IntersectWith(*herset)
	s3.DifferenceWith(*herset)
	s4.SymmetricDifference(*herset)
	fmt.Printf("myset:\t%v\t",myset)
	fmt.Printf("herset:\t%v\n",herset)
	w := tabwriter.NewWriter(os.Stdout,40,0,0,' ',0)
	fmt.Fprintf(w,"\033[1;34m%v\t%v\t%v\t%v\t\n\033[0m","Union","Intersection","Difference","SymmetricDifference")
	fmt.Fprintf(w,"%v\t%v\t%v\t%v\t\n",s1,s2,s3,s4)
	w.Flush()


}


func SetOperationOther_test(){

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

	var set1 = &IntSet{}
	var set2 =&IntSet{}
	var set3 =&IntSet{}
	set1.Add(uint64(5))
	set2.Add(uint64(74))
	set3.Add(uint64(1002))
	myset.AddAll(*set1,*set2,*set3)
	p("myset: %v\n", myset)
}
func main(){


	SetOperation_test()
	SetOperationOther_test()




}