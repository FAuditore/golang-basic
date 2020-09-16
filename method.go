package main

import "fmt"

func main() {
	/*
		方法:method 就是包含了接收者的函数，接收者可以是命名类型或结构体类型或指针
		func () 方法名(参数列表)(返回值列表){
		}
		方法名可以相同，接收者不同
	*/
	w1 := worker{
		name: "worker1",
		age:  1,
	}
	w1.work()
}

type worker struct {
	name string
	age  int
}

func (w worker) work() {
	fmt.Println(w, " working")
}
