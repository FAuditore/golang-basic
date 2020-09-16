package main

import "fmt"

func main() {
	/*
		 import –> const –> var –> init() –> main()
		常量表达式的值在编译期计算，而不是在运行期。
		每种常量的潜在类型都是基础类型:boolean、string 或数字。
	*/

	fmt.Println(100)
	fmt.Println("hello")

	//定义常量
	const PATH string = "/home/"

	//尝试修改 报错
	//PATH = "/home/hadoop"

	//定义一组
	const C1, C2, C3 = 100, 3.14, "haha"

	//如果初始值，默认和上一行一样
	const (
		a int = 100
		b
		c
	)
	fmt.Println(a, b, c)

	//枚举类型
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)
}
