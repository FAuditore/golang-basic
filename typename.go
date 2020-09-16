package main

import "fmt"

func main() {
	/*
		type 类型定义和别名
			1. type 类型名 Type  //定义一个新的类型
			2. type 类型名=Type	//给类型起别名
			给其他包的某些类型在该类中定义成一个新的类型，就可以添加/修改功能了
			起别名不能修改其他包的类型
	*/
	var i1 myint
	i1 = 200
	fmt.Printf("%T, %d\n", i1, i1) //main.myint 200

	res := fun1()
	fmt.Println(res(1, 2))

	var i3 myint2
	i3 = 100
	fmt.Printf("%T, %d\n", i3, i3) //int 100
}

type myint int

//定义函数类型
type myfun func(int, int) int

//返回值是一个自定义函数类型
func fun1() myfun {
	fun := func(a, b int) int {
		return a + b
	}
	return fun
}

//类型别名
type myint2 = int
