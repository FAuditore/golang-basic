package main

import (
	"fmt"
)

func main() {
	/*
		new(T) 返回一个指针，指向分配了零值的类型T  返回*T
		make(T,args) make只能创建slice,map,channel类型，返回一个有初始值的T类型(不是*T)

		匿名结构体
		name :=struct{
			field1 type1
			field2 type2
		}{
			field1:"",field2:123,
		}
	*/

	//定义结构体变量
	var p1 Person
	p1.name = "p1"
	p1.age = 18
	p1.sex = "nv"
	fmt.Println(p1)

	p2 := Person{}
	p2 = Person{"p2", 10, "nan"}
	fmt.Println(p2)

	p3 := Person{name: "p3", sex: "nv", age: 2}
	fmt.Println(p3)

	var pp1 *Person
	pp1 = &p1
	pp1.name = "pp1" //(*pp1).name="pp1"
	fmt.Println(p1)

	//new返回一个指向Person类型，赋了零值的Person指针
	pp2 := new(Person)
	fmt.Println(pp2) // &{0}
	pp2.name = "pp2"
	pp2.sex = "nan"
	pp2.age = 21
	fmt.Println(*pp2)

	//匿名结构体
	person2 := struct {
		name string
		age  int
	}{
		name: "person2", age: 15,
	}
	fmt.Println(person2)

	//匿名字段 不能有重复字段，默认把类型当作属性名
	w1 := Worker{"worker", 30}
	fmt.Println(w1)
	w1.string = "w1name"
	fmt.Println(w1)

	//结构体嵌套
	student := Student{}
	book1 := Book{"book", 500}
	student.name = "student"
	student.book = book1 //值传递
	fmt.Println(student)
	student.book.price = 300
	fmt.Println(book1)
	fmt.Println(student)
}

type Person struct {
	name string
	age  int
	sex  string
}

type Worker struct {
	string
	int
}

type Student struct {
	name string
	book Book
}
type Book struct {
	name  string
	price int
}
