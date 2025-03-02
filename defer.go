package main

import "fmt"

// defer特性
// 1. 多个defer之间肯定采用LIFO方式，
// 2. return最先执行写入返回值，然后执行defer收尾, 等收尾工作完成后，函数调用完成
// 3. defer不是在退出代码块的作用域后执行的，只是在当前函数退出之前被调用
// 4. 当外部变量作为函数参数传入到defer定义的函数中时，那么立即拷贝一份外部变量的值(新创建变量），这与闭包效果一样

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

// 参考特性3, 输出结果：
// leave code block
// returning function value
// defer called
func deferNotCalledAfterLeavingCodeBlock() {
	{
		defer fmt.Println("defer called")
		fmt.Println("leave code block")
	}

	fmt.Println("returning function value")
}

// 参考特性4，输出结果
// x 0xc000098040 1 0xc000098048 1
// y 0xc000098040 1 0xc000098058 0
func deferPutArgs() {
	i, j := 0, 0
	defer func(j int) { fmt.Println("y", &i, i, &j, j) }(j)
	i++
	j++
	fmt.Println("x", &i, i, &j, j)
}
