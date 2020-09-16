package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num1 := rand.Int()
	fmt.Println(num1)

	for i := 0; i < 5; i++ {
		fmt.Print(rand.Intn(5), " ") //rand.Intn(n) [0,n)
	}
	fmt.Println()

	rand.Seed(1) //同一个种子每次生成的随机数都一样
	num2 := rand.Intn(5)
	fmt.Print(num2)

	t1 := time.Now()
	fmt.Println(t1)
	timeStamp := t1.Unix() //以秒为单位时间戳
	fmt.Println(timeStamp)
	rand.Seed(timeStamp) //以时间戳为种子
	for i := 0; i < 5; i++ {
		fmt.Print(rand.Intn(5), " ")
	}
}
