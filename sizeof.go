package main

import (
	"fmt"
	"unsafe"
)

type IAnimal interface {
	Run()
}

type Human struct {
	Age int  //int为64位系统下为8字节
	Sex byte //byte为字节
}

type Person struct {
	Human
	Name string
}

func (p Person) Run() {

}

// 对不同切片，unsafe.Sizeof函数都返回24。其内存结构为指向数组的指针、长度和容量
// 对不同字符串，unsafe.Sizeof函数都返回16，这是因为string类型为一个结构体，含有两个域
// 一个域指向该字符串的指针，第二个域为字符串的长度，每个域为8个字节,
// 类似于C中的数组,里面已确定数组大小
// struct结构体需要考虑对齐和补齐, Human Age为8字节，Sex为1字节，考虑对齐和补齐，输出16字节
// 空结构体返回0, interface{}返回16字节,第一个类型指针, 第二个数据指针
func sizeof() {
	s := []int{1, 2, 3}
	s1 := []int{2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(s), unsafe.Sizeof(s1)) // 24, 24

	fmt.Println(unsafe.Sizeof("hel"), unsafe.Sizeof("helloworld")) // 16, 16

	a1 := [...]int{1, 2, 3}
	a2 := [...]int{1, 2, 3, 4}
	fmt.Println(unsafe.Sizeof(a1), unsafe.Sizeof(a2)) //24, 32

	fmt.Println(unsafe.Sizeof(Human{}), unsafe.Sizeof(Person{})) //16, 32
	var i IAnimal = Person{}
	fmt.Println(unsafe.Sizeof(struct{}{}), unsafe.Sizeof(i)) //0, 16
}
