package main

import (
	"fmt"
)

type LinkedList struct{
	value int64
	next *LinkedList

}

func (l *LinkedList)Sum() (sum int64){
	fmt.Println(l)
	//fmt.Println(l.value)
	if l.next == nil{
		sum = l.value
		return
	}
	sum =l.value+l.next.Sum()
	return



}

func (l *LinkedList)Len()(len int64){
	if  l.next == nil{
		return 1
	}
	len = 1 + l.next.Len()
	return


}

func (l *LinkedList)Average()int64{
	return l.Sum()/l.Len()



}


func (l *LinkedList)Max() (max int64){
	for l !=nil{
		if max < l.value {
			max = l.value
		}
		l = l.next
	}
	return
}



func (l *LinkedList)Min()(min int64){
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
	for i:=0;i<length-1;i++ {
		for j:=0;j<length-1-i;j++{
			if (l.next != nil)  && (l.value > l.next.value) {
				l.value, l.next.value = l.next.value, l.value
			}
			l = l.next
		}

	}



}

func (l *LinkedList)ReverseSort(){
	length := l.Len()
	temp := make([]int64,length)
	fmt.Printf(" Initial temp: %v\n",temp)
	l.Sort()
	for i := range temp{
		temp[i] = l.value
	}

	for i:=0;i<length/2;i++ {
		temp[i],temp[length-i] =  temp[length-i],temp[i]

	}

	for i:=0;i<length;i++{

		l.value = temp[i]
		l = l.next
	}

}

func (l *LinkedList)Remove(num int64) error{



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

func (l *LinkedList)Add(num int64){
	for l != nil{
		l = l.next
	}
	new := &LinkedList{num, nil}
	l = new

}

func (l *LinkedList)Has(num int64) (*LinkedList,bool){
	for l!=nil{
		if l.value == num{
			return l, true
		}
		l = l.next
	}
	return nil, false



}

func  (l *LinkedList)InSertBefore(target, num int64) (*LinkedList, error){
	p, has := l.Has(target)
	if !has{
		return nil,fmt.Errorf("no such elements whose value is %v\n", target)
	}
	if p == l {
		temp := p
		p := &LinkedList{num,temp}
		return p
	}
	for l.next != p {
		l = l.next
	}
	l.next = &LinkedList{num,p}
	return l.next




}

func  (l *LinkedList)InserAfter(target, num int64)(*LinkedList, error){
	p, has := l.Has(target)
	if !has {
		return nil,fmt.Errorf("no such elements whose value is %v\n", target)

	}
	temp := p.next
	p.next = &LinkedList{num,temp}
	return p.next,nil





}

func (l *LinkedList)InsertAs(index,num int64) (*LinkedList,error){
	if index > l.Len() || index <0{
		return nil, fmt.Errorf("Index:%v out of range\n",index)
	}
        if index == 0{
		temp := l
		l = &LinkedList{num,temp}
		return l, nil
	}

	for i:=0;l!=nil;i++{
		// get the previous pointer
		if i+1 == index{
			temp :=  l.next
			l.next = &LinkedList{num,temp}
			return l.next,nil
		}
		l = l.next
	}








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
	LinkedList_test()
}
