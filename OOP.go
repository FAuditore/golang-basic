package main

import "fmt"

func main() {
	/*
		结构体内匿名结构体的属性是该结构体的提升字段，可以直接访问
		重复的字段，会覆盖匿名结构体

		模拟继承
		type A struct{
			field
		}
		type B struct{
			A //匿名结构体
			field
		}

		模拟聚合关系 has a
		type A struct{
			field
		}
		type B struct{
			a A
			field
		}
	*/
	var student S
	student.P.name = "pname"
	student.name = "sname" //student.P.name="s" 提升字段
	student.age = 15
	student.grade = "2017"
	fmt.Println(student)

	c1 := cat{}
	c1.name = "cat"
	c1.like = "fish"
	c1.eat() //调用子类重写的父类方法 不重写默认调用父类的

}

type P struct {
	name string
	age  int
}

//P中的字段对于S来说就是提升字段，可以直接通过S进行访问
type S struct {
	P
	grade string
	name  string
}

type animal struct {
	name string
}

type cat struct {
	animal
	like string
}

func (a animal) eat() {
	fmt.Println(a, "eat")
}

func (c cat) eat() {
	fmt.Println(c, "eat fish")
}
