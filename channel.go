package main

import "fmt"

// 通道关闭
// 对一个关闭的通道执行接收操作先获取通道的值,通道的值获取完后，继续接受会得到对应类型的零值。
// 对一个关闭的通道再发送值就会导致panic。
// 关闭一个已经关闭的通道会导致panic。
// 输出结果：
// 0 true
// 0 false
// 0 false
func closeChannel() {
	ch := make(chan int, 1)
	ch <- 0
	close(ch)
	a, ok := <-ch
	fmt.Println(a, ok)
	a, ok = <-ch
	fmt.Println(a, ok)
	a, ok = <-ch
	fmt.Println(a, ok)

	// ch <- 1   // panic: send on closed channel
	// close(ch) // panic: close of closed channel
}

// 单向通道
// 通道默认为双通道，我们可以指定它在一个函数里面使用单向通道模式
func singleDirectionChannel() {
	done := make(chan struct{})
	// 输出通道，只能写
	producer := func(out chan<- int) {
		for i := 0; i < 5; i++ {
			out <- i
		}
		close(out)
	}
	// 输入通道，只能读
	consumer := func(in <-chan int) {
		for v := range in {
			fmt.Print(v)
		}
		done <- struct{}{}
	}
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)
	<-done
}
