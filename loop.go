package main

import "fmt"

// 两个问题：
// 在for循环和for-range循环中的变量，每次迭代后都会发生改变

// 输出结果:
// 0xc0000100a0
// 0xc0000100a8
func forVariableInLoop() {
	n := 2
	for i := 0; i < n; i++ {
		fmt.Println(&i)
	}
}

// 输出结果:
// 0xc000098040
// 0xc000098040
func forVariableOutOfLoop() {
	n := 2
	i := 0
	for i = 0; i < n; i++ {
		fmt.Println(&i)
	}
}

// 输出结果:
// 0xc000104040
// 0xc000104048
func forVariableInRangeLoop() {
	n := 2
	for i := range n {
		fmt.Println(&i)
	}
}

// 输出结果:
// 0xc0000100a0
// 0xc0000100a0
func forVariableOutOfRangeLoop() {
	n := 2
	i := 0
	for i = range n {
		fmt.Println(&i)
	}
}

func loop() {
	fmt.Println("while variable in for loop")
	forVariableInLoop()
	fmt.Println("while variable out of for loop")
	forVariableOutOfLoop()
	fmt.Println("while variable in for range loop")
	forVariableInRangeLoop()
	fmt.Println("while variable out of for range loop")
	forVariableOutOfRangeLoop()
}
