package main

import "fmt"

// 从1.22开始，在for循环和for-range循环中的变量，每次迭代后都会发生改变
// 之前则会保持一个变量
// 以下函数请在golang1.22之前和1.22之后分别进行测试

// 在1.22以前不允许使用for-range只能遍历可迭代类型,
// forVariableInRangeLoop以及forVariableOutOfRangeLoop在1.22之前会报错

func forVariableInLoop() {
	n := 2
	for i := 0; i < n; i++ {
		fmt.Println(&i)
	}
}

func forVariableOutOfLoop() {
	n := 2
	i := 0
	for i = 0; i < n; i++ {
		fmt.Println(&i)
	}
}

// 请在go1.22之后尝试
// func forVariableInRangeLoop() {
// 	n := 2
// 	for i := range n {
// 		fmt.Println(&i)
// 	}
// }

// func forVariableOutOfRangeLoop() {
// 	n := 2
// 	i := 0
// 	for i = range n {
// 		fmt.Println(&i)
// 	}
// }

func loop() {
	fmt.Println("while variable in for loop")
	forVariableInLoop()
	fmt.Println("while variable out of for loop")
	forVariableOutOfLoop()
	// 请在go1.22之后尝试
	// fmt.Println("while variable in for range loop")
	// forVariableInRangeLoop()
	// fmt.Println("while variable out of for range loop")
	// forVariableOutOfRangeLoop()
}
