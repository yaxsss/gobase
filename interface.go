package main

import (
	"fmt"
	"unsafe"
)

// 空接口类型，interface{}可接受任意类型的数据
// interface{}存有16个字节
// 输出结果
// 1 104 linux [1 2 3] map[1:2]
func emptyInterface() {
	anyway := func(args ...interface{}) {
		for _, arg := range args {
			fmt.Print(unsafe.Sizeof(arg), arg)
			fmt.Print(" ")
		}
		fmt.Println()
	}
	anyway(1, 'h', "linux", []int{1, 2, 3}, map[int]int{1: 2})
}

// 类型断言：
// arg.(type)可以获取类型结构
// arg.(struct{})用来Cast类型
func typeAssertion() {
	assertion := func(args ...interface{}) {
		for _, arg := range args {
			switch arg.(type) {
			case int:
				fmt.Print("int")
			case byte:
				fmt.Print("byte")
			case string:
				fmt.Print("string")
			case []int:
				fmt.Print("[]int")
			case map[int]int:
				fmt.Print("map[int]int")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
	assertion(1, 'h', "linux", []int{1, 2, 3}, map[int]int{1: 2})

	var s interface{} = struct{}{}
	if t, ok := s.(struct{}); ok {
		fmt.Println(t)
	}
}

// https://go.dev/wiki/InterfaceSlice
// var interfaceSlice []interface{} = dataSlice
// cannot use dataSlice(type []int) as type[] interface{} in assignment
// []interface{}为一个切片，指向的一个接口数组，其数组的元素类型与dataSlice的数组对象类型不一样
// 内存布局也就不一样，肯定不能直接转
func sliceToInterfaceArray() {
	var dataSlice []int = []int{1, 2, 3, 4, 6}
	var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
	for i, d := range dataSlice {
		interfaceSlice[i] = d
	}
}
