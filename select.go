package main

import (
	"fmt"
	"time"
)

// 每个case必须是操作channel，如读或写, 每个case上可以加上执行代码（如调用函数）,select开始就是对每个case代码求值，得到channel
// 多个channel可读或可写，select会随机公平地选出一个执行,其他而不会执行, 否则执行default子句
// channel可读的条件：不能为nil，channel必须有数据或者channel被关闭
// channel可写的条件：不能为nil，channel必须有空间
// 输出结果可能为：
//	read c2--->write c3--->read c1--->end
//	read c1--->read c2--->write c3--->end
//	write c3--->read c2--->read c1--->end

func selectCase() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	c3 := make(chan string, 1)
	c1 <- "read c1"
	c2 <- "read c2"
	v3 := "write c3"
	done := false
	for {
		select {
		case v1 := <-c1:
			fmt.Print(v1)
		case v2 := <-c2:
			fmt.Print(v2)
		case c3 <- v3:
			fmt.Print(v3)
		default:
			fmt.Print("end")
			done = true
		}
		if done {
			fmt.Println()
			break
		} else {
			fmt.Print("--->")
		}
	}
}

// 每个case上可以加上执行代码（如调用函数），select开始就是对每个case代码求值，得到channel
// 输出结果:
// <enter select><eval c1><read c1><leave select>
// <enter select><eval c1><end><leave select>
// 每次进入select都会调用runc1，输出eval c1
func selectCaseEval() {
	c1 := make(chan string, 1)
	c1 <- "<read c1>"
	runc1 := func() <-chan string {
		fmt.Print("<eval c1>")
		return c1
	}
	done := false
	for {
		fmt.Print("<enter select>")
		select {
		case v1 := <-runc1():
			fmt.Print(v1)
		default:
			fmt.Print("<end>")
			done = true
		}
		fmt.Println("<leave select>")
		if done {
			break
		}
	}
}

// donec被关闭，case <-donec每次都会执行，读出零值
// for循环每次执行，都是开启一个新的Timer，所以上面case <-time.After(time.Second)永远无法执行
// 执行结果：无任何输出，并陷入死循环
func selectCaseTimer() {
	donec := make(chan bool, 1)
	close(donec)
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("timer")
		case <-donec:
		}
	}
}
