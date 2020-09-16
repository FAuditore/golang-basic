package main

import (
	"fmt"
)

func main() {
	//iota 特殊常量，可以被编译器修改的常量 每次被使用时加1
	const (
		a = iota   //0
		b = iota   //1
		c          //2
		d = "haha" //iota = 3
		e          //"haha" iota = 4
		f = iota   //5
	)
	fmt.Println(a, b, c, d, e, f)

	const k = iota //0
	fmt.Println(k)
}
