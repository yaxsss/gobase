package main

// T类型的方法集由所有接收者类型为T的方法组成
// T指针类型的方法集由所有接受者类型*T或T的方法组成
// 接口类型的方法集由接口类型中的每种类型的方法集的交集
// 你不用同时定义接受类型为T和*T的函数，直接定义一个T的函数即可
// 接收者为T的方法，在调用时会发生拷贝，当T类型如果过大时，拷贝可能存在性能问题

type Window interface {
	Draw()
	SetParent()
}

type Label struct {
}

func (l Label) Draw() {
}

func (l *Label) SetParent() {
}

func CreateLabel() Window {
	// 无法编译通过, Label的方法集，只能是接受类型为Label的方法，没有SetParent，而Window接口需要你实现Window
	// cannot use Label{} (value of struct type Label) as Window value in variable declaration: Label does not implement Window (method SetParent has pointer receiver)
	// var w Window = Label{}
	var w Window = &Label{}
	return w
}
