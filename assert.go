package main

import "fmt"

func main() {
	/*
		接口断言
			一
			1.instance := 接口对象.(实际类型) //不安全
			2.instance,ok :=接口对象.(实际类型)

			二 switch
			switch instance :=接口对象.(type){
			case 实际类型1:
			case 实际类型2:
			...

			}
	*/

	c := cat2{"tom"}
	c.eat()
	allAnimalEat(c)
	allAnimalEat2(c)
}

type animal2 interface {
	eat()
}

type cat2 struct {
	name string
}

type dog2 struct {
	name string
}

func (c cat2) eat() {
	fmt.Println(c, " eat fish")
}

func (d dog2) eat() {
	fmt.Println(d, " eat meat")
}

func allAnimalEat(a2 animal2) {
	if ins, ok := a2.(cat2); ok {
		fmt.Println(ins)
		a2.eat()
	} else if ins, ok := a2.(dog2); ok {
		fmt.Println(ins)
		a2.eat()
	} else {
		fmt.Println("unkown")
	}
}

func allAnimalEat2(a2 animal2) {
	switch ins := a2.(type) {
	case cat2:
		fmt.Println(ins, " is a cat")
	case dog2:
		fmt.Println(ins, " is a dog")
	}
}
