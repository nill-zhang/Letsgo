package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Sortable struct {
	seq []int
}

func (s *Sortable) BubbleSort() {
	for i := 0; i < len(s.seq); i++ {
		// Start thinking from the inside loop
		// first we compare adjacent two and swap the two
		// if the first is greater than the second
		// after the first round the last one will be the greatest
		// then we just need to sort the previous len(s.sequence)-1
		// then len(s.sequence)-2......
		// why I don't start with j:=0, because I don't want to check
		// j+1 < len(s.sequence)-i
		for j := 1; j < len(s.seq)-i; j++ {
			if s.seq[j-1] > s.seq[j] {
				s.seq[j-1], s.seq[j] = s.seq[j], s.seq[j-1]
			}
		}
	}

}

func (s *Sortable) InsersionSort() {

	for i := 1; i < len(s.seq); i++ {
		for j := 0; j < i; j++ {

			// in all the elements before s.seq[i]
			// find the first element s[j] which is greater
			// than s.seq[i], insert s[i] before s[j]
			if s.seq[i] < s.seq[j] {
				temp := s.seq[i]
				// Note that I can not use
				// s.seq[j+1:i+1] = s.seq[j:i]
				// here for subslice assignment, single element
				// assignment is ok,but multiple elements
				// you need copy() or append()
				copy(s.seq[j+1:i+1], s.seq[j:i])
				s.seq[j] = temp
				break
			}
		}
	}
}

func (s *Sortable) BinarySort() {

	/*


		12 3 3 4 32 32  | 7 456 32 42 -234 5   swap(i,len(s)-1-i)
		  5 -234 3 4 32 7 | 32 456 32 42 3 12  BubbleSort(left) Min<--- BubbleSort(right) --->Max
		 - 234 5 2 34 3 7 | 32 32 42 3 12 456  swap(i,len(s)-1-i)
		       5 2 34 3 7 | 32 32 42 3 12      BubbleSort(left) Min<--- BubbleSort(right) --->Max
		       2 5 3 34 7 | 32 32 3 12 42      swap(i,len(s)-1-i)
		         5 3 32 7 | 32 34 3 12         BubbleSort(left) Min<--- BubbleSort(right) --->Max
		         3 5 7 32 | 32 3 12 34         ...
		           5 7 32 | 32 3 12            ...
		           5 3 32 | 32 7 12            ...
		           3 5 32 | 7 12 32            ...
		              5 7 | 32 12              ...
		              5 7 | 12 32              ...
		                7 | 12                 ...
	*/

	halfway := len(s.seq) / 2
	length := len(s.seq)
	flag := len(s.seq) % 2
	for j := 0; j < halfway; j++ {
		// compare symmetric elements, swap them to
		// make first half elments are smaller than the second half
		for i := j; i < halfway; i++ {
			if s.seq[i] > s.seq[length-1-i] {
				s.seq[i], s.seq[length-1-i] = s.seq[length-1-i], s.seq[i]
			}
		}
		// get the mininum from the first half
		for i := halfway - 1; i > j; i-- {

			if s.seq[i] < s.seq[i-1] {
				s.seq[i], s.seq[i-1] = s.seq[i-1], s.seq[i]
			}

		}

		// if s.seq has odd elements, symmetrically swap elements around it
		// after the smallest and greatest values are in its edge place,
		// compare the middle element with those two and take corresponding actions
		// if it is smaller than the smallest one or greater than the greatest one.
		if flag == 1 {

			// get the maximum from the second half
			for i := halfway + 2; i < len(s.seq)-j; i++ {

				if s.seq[i-1] > s.seq[i] {
					s.seq[i], s.seq[i-1] = s.seq[i-1], s.seq[i]
				}

			}

			if s.seq[halfway] < s.seq[j] {
				temp := s.seq[halfway]
				copy(s.seq[j+1:halfway+1], s.seq[j:halfway])
				s.seq[j] = temp
			}

			if s.seq[halfway] > s.seq[length-1-j] {
				temp := s.seq[halfway]
				copy(s.seq[halfway:length-2-j], s.seq[halfway+1:length-1-j])
				s.seq[length-1-j] = temp

			}

		}
		// get the maximum from the second half if flag == 0
		for i := halfway + 1; i < len(s.seq)-j; i++ {

			if s.seq[i-1] > s.seq[i] {
				s.seq[i], s.seq[i-1] = s.seq[i-1], s.seq[i]
			}

		}

	}

}

func (s *Sortable) ShaofengSort1(){
	for i:=0;i<len(s.seq);{
		flag := 0
		for j:=i+1;j<len(s.seq);j++{
			if s.seq[j] < s.seq[i]{
				s.seq[i],s.seq[j] = s.seq[j],s.seq[i]
				flag = 1
				break
			}


		}

		if flag == 1{
			continue
		}
		i++

	}
}


func (s *Sortable) ShaofengSort2(){
	for i:=0;i<len(s.seq);{
		flag := 0
		for j:=i+1;j<len(s.seq);j++{
			if s.seq[j] < s.seq[i]{
				temp := s.seq[j]
				copy(s.seq[i+1:j+1],s.seq[i:j])
				s.seq[i] = temp
				flag = 1
			}
		}

		if flag == 1{
			continue
		}
		i++

	}
}

func TimeSort(f func()) func() {
	start := time.Now()
	fmt.Printf("Entering %v\n", f)
	f()
	//stopf:= func(s time.Time) {
	//	fmt.Printf("leaving %v\n",f)
	//	fmt.Printf("Time spent %v\n",time.Since(s))
	//
	//}
	stopf := func() {
		fmt.Printf("leaving %v\n", f)
		fmt.Printf("Time spent %v\n", time.Since(start))

	}
	return stopf

}
func Sort_BenchMarkTest() {

	s1 := Sortable{[]int{}}
	for i := 0; i < 10000; i++ {
		s1.seq = append(s1.seq, rand.Intn(1000000))
	}
	s2 := Sortable{[]int{}}
	s3 := Sortable{[]int{}}
	s4 := Sortable{[]int{}}
	s5 := Sortable{[]int{}}
	s2.seq = append(s2.seq, s1.seq...)
	s3.seq = append(s3.seq, s1.seq...)
	s4.seq = append(s4.seq, s1.seq...)
	s5.seq = append(s5.seq, s1.seq...)

	//fmt.Printf("s1: %v\n", s1)
	//fmt.Printf("s2: %v\n", s2)
	//fmt.Printf("s3: %v \n", s3)
	//fmt.Printf("s4: %v\n", s4)
	//fmt.Printf("s5: %v\n", s5)
	TimeSort(s1.BinarySort)()

	TimeSort(s2.BubbleSort)()

	TimeSort(s3.InsersionSort)()

	TimeSort(s4.ShaofengSort1)()

	TimeSort(s5.ShaofengSort2)()
	//fmt.Printf("s1: %v\n", s1)
	//fmt.Printf("s2: %v\n", s2)
	//fmt.Printf("s3: %v\n", s3)
	//fmt.Printf("s4: %v\n", s4)
	//fmt.Printf("s5: %v\n", s5)

}

func main() {

	Sort_BenchMarkTest()

}
