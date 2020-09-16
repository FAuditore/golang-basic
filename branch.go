package main

import (
	"fmt"
)

func main() {

	num := 5
	if num > 20 {
		fmt.Println(">20")
	} else if num < 20 {
		fmt.Println("<20")
	} else {
		fmt.Println("=20")
	}

	if num2 := 10; num2 > 5 { //初始化仅if语句中可以使用num2
		fmt.Println("yes")
	}

	switch num { //自带break
	case 5:
		fmt.Println("5")
		fallthrough //运行完后继续运行下面的case,【无需匹配】
	case 10:
		fmt.Println("10")
		break //强制结束switch
		fmt.Println("10")
	default:
		fmt.Println("default")
	}

	switch number2 := 3; { //初始化仅该switch内部可见
	case number2 < 1:
		fmt.Println("number2<1")
	case number2 == 2, number2 == 3, number2 == 4:
		fmt.Println("number2==2/3/4")
	default:
		fmt.Println("number2>=5")
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	i := 0
	for i < 5 { //while(i<5)
		fmt.Printf("%d ", i)
		i++
	}

	for { //while(true)
		fmt.Print("again")
		break
	}

	fmt.Println()
out:
	for i := 0; i < 5; i++ { //out:循环标签名
		for j := 0; j < 5; j++ {
			if j == 3 {
				//break out1
				continue out
			}
			fmt.Printf("i:%d j:%d\n", i, j)
		}
	}
}
