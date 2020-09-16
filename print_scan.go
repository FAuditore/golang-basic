package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		%v 原样输出
		%T 打印类型
		%t bool类型
		%s 字符串
		%f 浮点
		%d 十进制整数
		%b 2进制整数
		%o 8进制
		%x %X 16进制 a-f A-F
		%c 字符
		%p 指针地址
	*/

	var x int
	var y float32
	fmt.Scanf("%d,%f", &x, &y)
	fmt.Printf("x:%d, y:%f\n", x, y)

	fmt.Println("输入一个字符串")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('a') //舍弃error 字符串结尾为a
	fmt.Println(s1)

	var s string
	fmt.Scanln(&s)
	fmt.Println("Scanln ", s)

}
