package main

import "fmt"

// 接口是一组行为的抽象

type Duck interface {
	swim()
	changeColor()
}

type ToyDuck struct {
	wet   bool
	color string
}

func (t ToyDuck) swim() {
	t.wet = true // 不改变t
}

func (t *ToyDuck) changeColor() {
	t.color = "green" // 会改变t
}

func NewDuck() Duck {
	return &ToyDuck{} // 接口类型必须返回指针
}

func main() {
	t := ToyDuck{wet: false, color: "yellow"}
	t0 := &ToyDuck{wet: false, color: "yellow"}
	t.swim()
	t0.swim()
	fmt.Printf("鸭子%v:指针鸭子：%v\n", t, t0) // 一样的
	t.changeColor()
	t0.changeColor()
	fmt.Printf("鸭子%v:指针鸭子：%v\n", t, t0) // 一样的
}

type A ToyDuck   // new type (没有Duck的方法）
type B = ToyDuck // alias

var duck4 ToyDuck
var duck5 *ToyDuck

// 指针

type Node struct {
	// 自引用只能使用指针 ==> 否则无法计算内存大小
	// left Node
	// right Node
	left  *Node
	right *Node
}
