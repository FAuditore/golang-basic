package main

import "fmt"

func main() {
	/*
		数组 var 数组名 [长度]数据类型
			 var 数组名 = [长度]数据类型{元素1,元素2...}
				 数组名 := []数据类型{元素1,元素2...}
	*/
	var arr1 [10]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 4
	arr1[3] = 3
	fmt.Println("arr1 ", arr1[1])
	fmt.Println(len(arr1)) //10 长度
	fmt.Println(cap(arr1)) //10 容量

	var arr2 = [4]int{}
	fmt.Println("arr2 ", arr2)

	var arr3 = []string{"a", "b", "c", "bcd"}
	fmt.Println("arr3 ", arr3)

	arr4 := [...]string{"a", "b"}
	fmt.Println("arr4 ", arr4)

	arr5 := []int{0: 1, 3: 5} //省略...变slice
	fmt.Printf("arr5 %T %v\n", arr5, arr5)

	for index, value := range arr1 {
		fmt.Print("(", index, value, ") ")
	}

	/*
		值类型：拷贝传递的是副本
			int,float,string,bool,array
		引用类型：存储的是数据的地址 默认值是nil
			slice,map,function,pointer,channel...
	*/

	array1 := [...]int{1, 2, 3}
	array2 := [...]int{1, 2, 3}
	fmt.Println(array2 == array1) //true 且必须长度一样

	a2 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7}, {10}}
	fmt.Println(len(a2))
	fmt.Println(cap(a2))
	fmt.Println(a2)

	//遍历二维数组
	for _, arr := range a2 {
		for _, value := range arr {
			fmt.Print(value, " ")
		}
	}
}
