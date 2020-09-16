package main

import "fmt"

func main() {
	/*
		数组指针  *[n]type  一个指针，存储数组的地址
		指针数组  [n]*type	一个数组，存储的数据类型是指针

		函数指针 一个指向函数的指针
			slice,map,function
		指针函数 一个函数，返回值是一个指针

		指针之间==,只有当它们指向同一个变量或全部是nil时才相等。

		不可寻址: 即不可调用*T方法
		字符串中的字节；
		map 对象中的元素（slice 对象中的元素是可寻址的，slice的底层是数组）；
		常量；
		包级别的函数等
	*/

	a := 10
	fmt.Printf("a的地址: %p\n", &a)
	var p1 *int
	p1 = &a         // p, of type *int, points to a
	fmt.Println(p1) //a的地址
	fmt.Printf("p1的地址: %p\n", &p1)
	fmt.Println(*p1)

	*p1 = 200
	fmt.Println(a) //200

	var p2 **int
	p2 = &p1
	fmt.Printf("%T, %T, %T\n", a, p1, p2)
	fmt.Println(p2) //p1的地址

	//数组指针
	arr1 := [4]int{1, 2, 3, 4}
	p3 := &arr1
	fmt.Println(p3)               //&[1,2,3,4]
	fmt.Printf("%T %p\n", p3, p3) //*[4]int 数组的地址

	p3[0] = 200 // (*p3)[0]=200
	fmt.Println(arr1)

}
