package main

import "fmt"

// 闭包是一个值（为函数），它可以访问定义其函数体外的变量（引用方式）
// 延迟绑定：闭包捕获的时机，不是在定义时被捕获，而是在调用时采取获取当前作用域的变量值
// 当定义闭包时有传入参数，那么这个参数会立即发生拷贝

// 输出结果：
// 同样受到loop变量的影响（参考loop.go), i已经是新变量，引用的是新变量
// 0xc00001006c 0xc00009a000
// 0xc00001006c 0
// 0xc00009a000 1
func lateBinding() {
	l := [](func() int32){}
	m := int32(2)
	for i := range m {
		fmt.Print(&i, " ")
		l = append(l, func() int32 { fmt.Print(&i, " "); return i })
	}
	fmt.Println()
	for _, x := range l {
		k := x()
		fmt.Println(k)
	}
}
