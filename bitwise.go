package main

import "fmt"

func main() {
	/*
		按位& 都为1则为1，有一个0为0
		按位| 都为0则为0，有一个1为1
		异或^
			二元a^b: 不同为1，相同为0
			一元^a: 按位取反 1->0 0->1
		位清空&^
			二元a&^b: 对于b的每个数值，如果为0，则取a对应位的值 如果为1，就取0  ->a&(^b)
			一元&^a:
	*/

	a := 60 //0011 1100
	b := 13 //0000 1101
	fmt.Printf("a: %d, %b\n", a, a)
	fmt.Printf("b: %d, %b\n", b, b)
	fmt.Printf("^a: %d, %b\n", ^a, ^a)
	fmt.Printf("a&b: %d, %b\n", a&b, a&b)
	fmt.Printf("a|b: %d, %b\n", a|b, a|b)
	fmt.Printf("a^b: %d, %b\n", a^b, a^b)
	fmt.Printf("a&^b: %d, %b\n", a&^b, a&^b)
}
