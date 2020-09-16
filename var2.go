package main

import "fmt"

func main() {

	var num int
	fmt.Printf("%d, %p\n", num, &num)

	var str string
	fmt.Println(str)

	//静态语言，变量类型和赋值类型必须一致
	var name = "abc"
	//name = 100
	fmt.Print(name)
}
