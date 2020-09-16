package main

import (
	"fmt"
)

func main() {
	/*
		 map 无序键值对
		 	 键不能重复
			 map[key类型][value类型]
		     引用类型
	*/
	//key是int value是string 没有初始化 map1==nil
	var map1 map[int]string
	fmt.Println(map1 == nil)
	//创建
	map1 = make(map[int]string)

	var map2 = make(map[int]string)
	map3 := map[string]int{"a": 3, "b": 5}
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)

	map2[1] = "hello"
	map2[2] = "aaaa"
	map2[3] = "bbbb"
	map2[100] = "dddd"
	fmt.Println(map2)

	//ok是bool类型，代表key是否存在
	value, ok := map2[4]
	fmt.Println(value, ok)

	delete(map2, 3)
	fmt.Println(map2)

	//遍历map
	for key, value := range map2 {
		fmt.Println(key, value)
	}

	slice1 := make([]map[int]string, 0, 3)
	slice1 = append(slice1, map2)
	fmt.Println(slice1[0])

	p1 := pointerFunc()
	fmt.Println(p1)  //&[1,2,3,4]
	fmt.Println(*p1) //[1,2,3,4]

	//坑：数组和map的遍历，range传的是值拷贝，value的地址不变
	structs := []student{
		{name: "abc", age: 10},
		{name: "bbb", age: 100},
	}

	for _, v := range structs {
		fmt.Println(v)
		fmt.Printf("%p\n", &v)
		fmt.Println("----------")
	}
	fmt.Printf("%p\n", &structs[0])
	fmt.Printf("%p\n", &structs[1])
}

//指针函数
func pointerFunc() *[4]int {
	arr := [4]int{1, 2, 3, 4}
	return &arr
}
