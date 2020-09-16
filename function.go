package main

import "fmt"

func main() {
	/*
		 func funcName(param1 type1, param2 type2...) (output1 type1,output2,type2...){
			...
			return value1,value2
		}
		defer:
		1.defer 延迟执行 在外围函数执行完之后执行
		func ...{
		file.Open("file")
		defer file.close()
		}
		2.多个defer后进先出 栈
		3.defer发生函数传递时，调用时已经传递参数了，只是暂时不执行函数中的代码
		4.在return语句之前执行

		匿名函数
		1.做参数，回调函数
		2.做返回值，形成闭包结构

		高阶函数
		接收了函数做参数的函数
		回调函数
		作为另一个函数的参数的函数

		闭包
		一个外层函数，有内层函数操作外层函数的局部变量，并且返回值为内层函数
		则内层函数和外层函数的局部变量，统称为闭包结构
		局部变量生命周期会发生改变(不会销毁)


		函数指针 一个指向函数的指针 var a = func(a,b int)int{
									return a+b
									}
			slice,map,function
		指针函数 一个函数，返回值是一个指针
	*/
	i, _ := getSum(1, 2)
	fmt.Println(i)

	i = getSum2(1, 2, 3, 4, 5, 6)
	fmt.Println(i)

	s1 := []int{1, 2, 3, 4}
	i = getSum2(s1...)
	fmt.Println(i)

	defer p("step1")
	fmt.Println("step2")
	s := "step3"
	defer p(s) //调用时已经传递参数了，只是没执行 结果为step3
	s = "ssssss"
	fmt.Println("step4")

	fmt.Println(p) //函数名对应函数体的地址
	var c func(string)
	fmt.Println(c) //<nil>
	c = p
	fmt.Println(c)
	c("123")

	//匿名函数
	func() {
		fmt.Println("匿名函数")
	}()

	fn := func() {
		fmt.Println("我也是匿名函数")
	}
	fn()

	//带参匿名函数
	func(a, b int) {
		fmt.Println(a + b)
	}(1, 2)

	//匿名函数调用结果给res
	res := func(a, b int) int {
		return a + b
	}(10, 20)
	fmt.Println(res)

	//匿名函数赋值给res2
	res2 := func(a, b int) int {
		return a + b
	}
	res2(1, 2)

	//高阶函数
	res3 := oper(1, 2, add)
	fmt.Println(res3)

	res5 := oper(3, 4, func(i int, i2 int) int {
		return i * i2
	})
	fmt.Println(res5)

	v1 := increment() //闭包
	fmt.Println(v1()) //1
	fmt.Println(v1()) //2
}

//返回多个值(type1,type2) 也可以(name1 type1,name2 type2) 函数中可以使用name1和name2
func getSum(a int, b int) (int, result int) {
	return a + b, a - b
}

//arg为int切片
func getSum2(arg ...int) int {
	sum := 0
	for _, value := range arg {
		sum += value
	}
	return sum
}

func p(s string) {
	fmt.Println(s)
}

//回调函数
func add(a, b int) int {
	return a + b
}

//高阶函数
func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun)
	return fun(a, b)
}

//闭包
func increment() func() int {
	i := 0 //若在内层函数中调用了，则不会销毁
	//定义一个匿名函数并返回
	fun := func() int {
		i++
		return i
	}
	return fun
}
