package main

import "fmt"

func main() {
	/*
		基本类型 数字 字符串 布尔型
		复合类型 数组 结构体 slice map

		值类型：默认拷贝传递的是副本
			int,float,string,bool,array,struct
		引用类型：存储的是数据的地址 默认值是nil
			slice,map,function,pointer,channel...

		v,ok:=map[key]
		v,ok:=<-chan
		v,ok:=value.(T)
	*/

	//布尔类型
	var b1 bool
	b1 = true
	b2 := false
	fmt.Printf("%T, %t\n%T, %t\n", b1, b1, b2, b2)

	//整型 64位系统int/uint默认64位 -2^63-2^63-1 int不等同于int64
	//byte==uint8  rune==int32
	var i1 int8
	i1 = 127
	fmt.Println(i1)
	var i2 rune
	i2 = 250
	fmt.Println(i2)

	//浮点 保留自动四舍五入 默认float64
	var f1 float32
	f1 = 1.23456
	fmt.Printf("%T, %.3f\n", f1, f1)

	//字符串
	//字符串为只读类型 可以str[n]读取 但不可以修改值
	s1 := "刘波"
	fmt.Printf("%T, %s\n", s1, s1)

	v1 := 'A' //单引号为rune类型
	v2 := "A"
	fmt.Printf("%T, %d,%c,%q\n", v1, v1, v1, v1)
	fmt.Printf("%T, %s\n", v2, v2)

	//类型转换 go为静态语言，要求赋值运算定义必须类型一致
	var a1 float32
	a1 = 10.9569
	var b int16
	b = int16(a1) //10
	fmt.Println(a1, b)

	var a float64
	fmt.Println(1/a, -1/a, a/a) //+Inf,-Inf,NaN
	var x complex128
	x = 1 + 2i
	fmt.Println(x * x)

}

type student struct {
	name string
	age  int
}
