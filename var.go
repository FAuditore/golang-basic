package main

import "fmt"

func main() {
	/*
		基本类型：整型，浮点型，布尔，字符串
		复合类型：array,slice,map,struct,pointer,function,channel...
	*/
	//定义之后赋值
	var age int
	age = 18
	fmt.Printf("age = %d\n", age)

	//类型推断
	var number = 15
	fmt.Printf("number = %d\n", number)

	var string = "abc"
	fmt.Printf("type = %T, value = %s\n", string, string)

	//简短声明
	sum := 100
	fmt.Printf("sum = %d\n", sum)

	//同时定义
	var a, b, c = 1, 2.5, "ccc"
	fmt.Println(a, b, c)

	//集合
	var (
		name = "abc"
		sex  = "男"
	)
	fmt.Println(name, sex)
}
