package main

import (
	"errors"
	"fmt"
	"net"
	"os"
)

func main() {
	/*
		error: 内置数据类型，内置的接口
			Error() string

		errors.New(),创建一个error
		fmt.Errorf(),创建一个error

	*/
	file, err := os.Open("test.txt")
	if err != nil {
		//log.Fatal(err) = print() + return
		fmt.Println(err)
		if ins, ok := err.(*os.PathError); ok {
			fmt.Println("Op ", ins.Op)
			fmt.Println("Err ", ins.Err)
			fmt.Println("Path", ins.Path)
			return
		}
	}
	fmt.Println(file, " OPEN SUCCESS")

	//errors.New创建error
	err1 := errors.New("this is an error")
	fmt.Println(err1)

	//通过fmt.Errorf创建error
	err2 := fmt.Errorf("error code: %d", 100)
	fmt.Println(err2)

	err3 := checkAge(-1)
	if err3 != nil {
		fmt.Println(err3)
	}

	addr, err4 := net.LookupHost("www.basdfafadsaidu.com")
	if ins, ok := err4.(*net.DNSError); ok {
		fmt.Println(ins.Err)        //no such host
		fmt.Println(ins.Name)       //www.basdfafadsaidu.com
		fmt.Println(ins.Server)     //""
		fmt.Println(ins.IsNotFound) //true
	}
	fmt.Println(addr) //[]

	str, err := get(-5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)

}

func checkAge(age int) error {
	if age < 0 {
		//return errors.New("年龄不合法")
		err := fmt.Errorf("您的输入有误: %d", age)
		return err
	}
	fmt.Println(age)
	return nil

}

type myError struct {
	name string
	age  int
}

//实现error接口
func (me *myError) Error() string {
	return fmt.Sprintf("%s %d ", me.name, me.age)
}

func get(age int) (string, error) {
	if age < 0 {
		return "", &myError{name: "error:age<0 ", age: age}
	}
	return "ok", nil
}
