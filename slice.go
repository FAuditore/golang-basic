package main

import "fmt"

func main() {
	/*
		切片存储的是一个数组的引用
		如果oldCap*2<Cap newCap = Cap
		否则
			oldLen<1024 newCap=2*oldCap
			oldLen>=1024 newCap=oldCap*1.25
	*/
	s1 := make([]int, 0, 3)
	fmt.Println(s1)
	fmt.Println(len(s1), " ", cap(s1))
	fmt.Printf("%p\n", s1)

	s1 = append(s1, 1, 2, 3, 4)
	fmt.Println(len(s1), " ", cap(s1))
	fmt.Printf("%p\n", s1)

	arr := [4]int{1, 2, 3, 4}
	s2 := arr[:3] //从arr[0]到arr[2]
	s3 := arr[2:]
	fmt.Println(s2)
	fmt.Println(s3)

	arr1 := [3]int{}
	arr2 := arr1 //数组是值类型，拷贝类型是值拷贝即深拷贝
	fmt.Printf("%p\n", &arr1)
	fmt.Printf("%p\n", &arr2)

	s4 := s1 //切片是引用类型，拷贝类型为浅拷贝
	fmt.Printf("%p\n", s4)
	fmt.Printf("%p\n", s1)

}
