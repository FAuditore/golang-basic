package main

import "fmt"

func main() {
	/*
		panic() 进入恐慌，中断程序执行 panic上面的代码会执行，panic后面的不执行
		recover() 让程序回复，在defer中执行
	*/
	defer fmt.Println("defer main 1")
	funA()
	defer fmt.Println("defer main 2")
}

func funA() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg, " 程序恢复") //恢复之后对主函数没有影响了 msg就是panic的参数
		}
	}()
	defer fmt.Println("defer A 1")
	for i := 0; i < 10; i++ {
		if i == 5 {
			panic("A panic")
		}
		fmt.Println(i)
	}
	defer fmt.Println("defer A 2")
}
