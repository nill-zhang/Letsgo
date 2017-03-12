package main

import (
	"fmt"
	"strconv"
	"strings"
)

type LinkedList struct{
	value int
	next *LinkedList

}

func (l *LinkedList)Sum() (sum int){
	if l.next == nil{
		sum = l.value
		return
	}
	sum =l.value+l.next.Sum()
	return



}

func (l *LinkedList)Len()(len int){
	if  l.next == nil{
		return 1
	}
	len = 1 + l.next.Len()
	return


}

func (l *LinkedList)Average()int{
	return l.Sum()/l.Len()



}


func (l *LinkedList)Max() (max int){
	for l !=nil{
		if max < l.value {
			max = l.value
		}
		l = l.next
	}
	return
}



func (l *LinkedList)Min()(min int){
	for l !=nil{
		if min < l.value {
			min = l.value
		}
		l = l.next
	}
	return



}


func (l *LinkedList)Sort(){
	length := l.Len()
	// keep the head pointer
	beg := l
	for i:=0;i<length-1;i++ {
		// start over from the head
		l := beg
		for j:=0;j<length-1-i ;j++{
			if l.next != nil {
				if (l.next != nil)  && (l.value > l.next.value) {

					temp := l.value
					l.value = l.next.value
					l.next.value = temp

				}
				l = l.next
			}

		}

	}



}

func (l *LinkedList)Elem() []int{
	length := l.Len()
	temp := make([]int,length)
	for i := range temp{
		temp[i] = l.value
		l = l.next
	}
	return temp




}

func (l *LinkedList)ReverseSort(){
	l.Sort()
        temp := l.Elem()
	length := len(temp)
	for i:=0;i<length/2;i++ {
		temp[i],temp[length-i] =  temp[length-i],temp[i]

	}

	for i:=0;i<length;i++{

		l.value = temp[i]
		l = l.next
	}

}

func (l *LinkedList)Remove(num int) error{



	for l.next != nil{
		// check the first element, if it is
		// the element we are looking for
		// return
		if l.value == num {
			l = l.next
			return nil
		}
		if l.next.value == num{
			// don't need to check l.next.next != nil
			// if its nil we delete the last element
			l.next = l.next.next
			return nil

		}
		l = l.next
	}
	return fmt.Errorf("did not find the element with value %v\n",num)


}

func (l *LinkedList)Clear(){
	for l!=nil{
		l.value = 0
		l = l.next
	}



}

func (l *LinkedList)Add(num int){
	for l.next != nil{
		l = l.next
	}
	new := &LinkedList{num, nil}
	l.next = new

}

func (l *LinkedList)Has(num int) (*LinkedList,bool){
	for l!=nil{
		if l.value == num{
			return l, true
		}
		l = l.next
	}
	return nil, false



}

func  (l *LinkedList)InsertBefore(target, num int) (*LinkedList, error){
	p, has := l.Has(target)
	if !has{
		return nil,fmt.Errorf("no such elements whose value is %v\n", target)
	}
	if p == l {
		temp := p
		p := &LinkedList{num,temp}
		return p,nil
	}
	for l.next != p {
		l = l.next
	}
	l.next = &LinkedList{num,p}
	return l.next,nil




}

func  (l *LinkedList)InsertAfter(target, num int)(*LinkedList, error){
	p, has := l.Has(target)
	if !has {
		return nil,fmt.Errorf("no such elements whose value is %v\n", target)

	}
	temp := p.next
	p.next = &LinkedList{num,temp}
	return p.next,nil





}

func (l *LinkedList)InsertAs(index,num int)(*LinkedList,error) {
	if index > l.Len() || index <0{
		return nil, fmt.Errorf("Index:%v out of range\n",index)
	}
        if index == 0{
		originalValue := l.value
		temp := l.next

		originalHead := &LinkedList{originalValue,temp}
		l.value = num
		l.next = originalHead
		return l, nil
	}

	for i:=0;l!=nil;i++{
		// get the previous pointer
		if i+1 == index{
			temp :=  l.next
			l.next = &LinkedList{num,temp}
			break
		}
		l = l.next
	}
	return l.next,nil
}


func (l *LinkedList)String()string{
	elements := l.Elem()
	s :=[]string{}
	for i:=0;i<len(elements);i++{
		s = append(s,strconv.Itoa(elements[i]))
	}
	return "["+strings.Join(s,",")+"]"






}
func LinkedListInclusive_test(){
	p := &LinkedList{0,nil}
	p.Add(-6)
	p.Add(54)
	p.Add(17)
	p.Add(69)
	p.Add(16)
	_, has := p.Has(6)
	fmt.Printf("has 6 ? %v\n",has)
	_, has = p.Has(-4)
	fmt.Printf("has -4 ? %v\n",has)
	_, has = p.Has(67)
	fmt.Printf("has 67 ? %v\n",has)
	_, has = p.Has(69)
	fmt.Printf("has 69 ? %v\n",has)
	_, has = p.Has(-16)
	fmt.Printf("has -16 ? %v\n",has)
	p.Remove(69)
	_, has = p.Has(69)
	fmt.Printf("has 69 ? %v\n",has)
	fmt.Printf("Average: %v\n", p.Average())
	fmt.Printf("Sum: %v\n", p.Sum())
	fmt.Printf("Max: %v\n", p.Max())
	fmt.Printf("Min: %v\n", p.Min())
	fmt.Printf("p: %s\tlength: %d\n", p,p.Len())
	p.Sort()
	fmt.Printf("p: %s\n", p)
	p.InsertAs(0,777)
	p.InsertAs(0,512)
	p.InsertAs(3,43)
	p.InsertAs(3,90)
	p.InsertAs(5,-30)
	p.InsertBefore(777,76)
	p.InsertAfter(512,13)
	fmt.Printf("p: %s\n", p)``
	p.Sort()
	fmt.Printf("p: %s\n", p)
	p.Clear()
	fmt.Printf("p: %s\n", p)



}

func LinkedList_test(){
	var p *LinkedList
	fmt.Printf("\nType : %T, Value(p): %[1]v\n",p)
	fmt.Printf("Type: %T, Value(*): %[1]v\n",((*LinkedList)(nil)))
	fmt.Printf("Type: %T, Value(new): %[1]v\n",new(LinkedList))

	temp := LinkedList{-4,p}
	fmt.Printf("Type : %T, value(&): %[1]v\n",&LinkedList{-4,p})
	fmt.Printf("Type : %T, value(variable): %[1]v\n",&temp)

	a := LinkedList{3,&LinkedList{3,&LinkedList{4,&LinkedList{5,&LinkedList{6,&LinkedList{7,
		&LinkedList{10,&LinkedList{40,&LinkedList{98,&LinkedList{-4,p}}}}}}}}}}
	b := LinkedList{3,&LinkedList{3,&LinkedList{4,&LinkedList{5,&LinkedList{6,&LinkedList{7,
		&LinkedList{10,&LinkedList{40,&LinkedList{98,&LinkedList{-4,(*LinkedList)(nil)}}}}}}}}}}
	c := LinkedList{3,&LinkedList{3,&LinkedList{4,&LinkedList{5,&LinkedList{6,&LinkedList{7,
		&LinkedList{10,&LinkedList{40,&LinkedList{98,&LinkedList{-4,new(LinkedList)}}}}}}}}}}
	//a := LinkedList{98,&LinkedList{-4,p}}

	fmt.Printf("a:Sum(%#v):%v\n",a,a.Sum())
	fmt.Printf("b:Sum(%#v):%v\n",b,b.Sum())
	fmt.Printf("c:Sum(%#v):%v\n",c,c.Sum())



}

func main() {
	//LinkedList_test()
	LinkedListInclusive_test()
}
