package main

import "fmt"

// 多个defer之间肯定采用LIFO方式，
// return最先执行写入返回值，然后执行defer收尾, 等收尾工作完成后，函数调用完成
// defer不是在退出代码块的作用域后执行的，只是在当前函数退出之前被调用
// 当外部变量作为函数参数传入到defer定义的函数中时，那么立即拷贝一份外部变量的值，这与闭包效果一样

func deferOrder() int {
	var i int
	defer func() {
		i++
		fmt.Println("the first defer", i)
	}()

	defer func() {
		i++
		fmt.Println("the second defer", i)
	}()
	i++
	return i
}
