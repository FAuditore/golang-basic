package main

import "fmt"

func main() {
	/*
		接口 一组方法签名 非嵌入式

		当需要接口类型对象时，可以使用任意实现类代替 testInterface(value interfaceName)
		接口对象不能访问实现类中的属性

		一个接口的实现 可以看做实现本身的类型，访问实现类的属性和方法
					 可以看作接口类型，只能访问接口中的方法,不能访问实现类的属性和方法
					 一个类实现了哪个接口的所有方法，才视为实现了这个接口
		接口用法：
				1.一个函数接受接口类型的参数，那么可以传入任意实现类调用
				2.定义一个类型为接口类型，赋值为实现类的对象

		空接口 不包含任何方法，任何类型都可视为实现了这个接口

		var _ interface = name 	确保name变量实现了interface接口
		var _ eat = animal

		type iface struct {
			tab  *itab
			data unsafe.Pointer
		}

		type itab struct {
			inter *interfacetype
			_type *_type
			hash  uint32 // copy of _type.hash. Used for type switches.
			_     [4]byte
			fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
		}


		type eface struct {
			_type *_type
			data  unsafe.Pointer
		}

		type _type struct {
		size       uintptr
		ptrdata    uintptr // size of memory prefix holding all pointers
		hash       uint32
		tflag      tflag
		align      uint8
		fieldAlign uint8
		kind       uint8
		// function for comparing objects of this type
		// (ptr to object A, ptr to object B) -> ==?
		equal func(unsafe.Pointer, unsafe.Pointer) bool
		// gcdata stores the GC type data for the garbage collector.
		// If the KindGCProg bit is set in kind, gcdata is a GC program.
		// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
		gcdata    *byte
		str       nameOff
		ptrToThis typeOff
	}
	*/
	var _ USB = mouse{}

	m1 := mouse{
		name: "mouse1",
	}
	testInterface(m1)

	var usb1 USB
	usb1 = m1
	fmt.Printf("%T, %v\n", usb1, usb1)
	//usb1.name接口不能调用实现类的属性，只能调用方法
	usb1.start() //自动调用m1的实现

	//空接口类型
	var m2 a = m1
	testa(m2)
	fmt.Println(m2)

	//可以存储任意类型的map
	map1 := make(map[string]interface{})
	map1["name"] = "mapname"
	map1["age"] = 18
	testa(map1)

	m3 := mouse{"use usb"} //是UUSB接口类型，能调用USB接口方法
	fmt.Println(m3)

}

type USB interface {
	start()
	end()
}
type a interface {
}

type UUSB interface {
	USB
	connect()
}

//实现类
type mouse struct {
	name string
}

func (m mouse) start() {
	fmt.Println(m, " start")
}

func (m mouse) end() {
	fmt.Println(m, " end")
}

func (m mouse) connect() {
	fmt.Println(m, " connect")
}

func testInterface(usb USB) {
	usb.start()
	usb.end()
}

//接受任意类型的参数
func testa(a interface{}) {
	fmt.Println(a)
}
